package model

import (
	"fmt"
	"strconv"
	"strings"
)

// FieldDesc 定义了全部的数据类型 以及对类型的解析
type FieldType int

const (
	FieldType_None   FieldType = 0
	FieldType_Int32  FieldType = 1
	FieldType_Int64  FieldType = 2
	FieldType_UInt32 FieldType = 3
	FieldType_UInt64 FieldType = 4
	FieldType_Float  FieldType = 5
	FieldType_String FieldType = 6
	FieldType_Bool   FieldType = 7
	FieldType_Enum   FieldType = 8
	FieldType_Struct FieldType = 9
	FieldType_Table  FieldType = 10 // 表格, 仅限二进制使用
	FieldType_Int    FieldType = 11 // = int64
	FieldType_Text   FieldType = 12 // 可以有中文逗号
	// FieldType_Vector3 FieldType = 13
	// FieldType_Vector2    FieldType = 14
	FieldType_Key        FieldType = 15 // key 这里会检查是否重复
	FieldType_CustomEnum FieldType = 16 // 类型行定义的枚举类型
	FieldType_Dict       FieldType = 17 // 字典类型 [type]=[type]
	FieldType_Double     FieldType = 18 // double类型 同float
)

// 一列的描述
type FieldDescriptor struct {
	Name string

	Type FieldType

	Complex *Descriptor // 复杂类型: 枚举或者结构体

	Order int32 // 在Descriptor中的顺序

	Meta *MetaInfo // 扩展字段

	IsRepeated bool

	Is2DArray bool

	EnumValue int32 // 枚举值

	Comment string // 注释

	ExportClient bool // 客户端是否导出

	ExportServer bool // 服务器是否导出

	IsKey bool // 是否为key

	Parent *Descriptor

	EnumMap map[string]int

	IsCustomEnum bool

	DictTypes [2]FieldType // 字典类型的种类
}

func NewFieldDescriptor() *FieldDescriptor {
	return &FieldDescriptor{
		Meta: NewMetaInfo(),
	}
}

func (self *FieldDescriptor) Tag() int32 {

	return MakeTag(int32(self.Type), self.Order)
}

func MakeTag(t int32, order int32) int32 {
	return t<<16 | order
}

func (self *FieldDescriptor) Equal(fd *FieldDescriptor) bool {

	if self.Name != fd.Name {
		return false
	}

	if self.Type != fd.Type {
		return false
	}

	if self.Meta.String() != fd.Meta.String() {
		return false
	}

	if self.IsRepeated != fd.IsRepeated {
		return false
	}

	if self.EnumValue != fd.EnumValue {
		return false
	}

	if self.complexName() != fd.complexName() {
		return false
	}

	return true
}

func (self *FieldDescriptor) complexName() string {
	if self.Complex != nil {
		return self.Complex.Name
	}

	return ""
}

// 自动适配结构体和枚举输出合适的类型, 类型名为go only
func (self *FieldDescriptor) TypeString() string {
	if self.Complex != nil {
		return self.Complex.Name
	} else {
		return FieldTypeToString(self.Type)
	}
}

func (self *FieldDescriptor) KindString() string {
	return FieldTypeToString(self.Type)
}

func (self *FieldDescriptor) String() string {

	var repString string
	if self.IsRepeated {
		repString = "repeated "
	}

	return fmt.Sprintf("name: '%s' %stype: '%s'", self.Name, repString, self.TypeString())
}

func (self *FieldDescriptor) DefaultValue() string {

	if v := self.Meta.GetString("Default"); v != "" {
		return v
	}

	switch self.Type {
	case FieldType_Int32,
		FieldType_UInt32,
		FieldType_Int64,
		FieldType_UInt64,
		FieldType_Int,
		FieldType_Float,
		FieldType_Double:
		return "0"
	case FieldType_Bool:
		return "false"
	case FieldType_Enum:

		if self.Complex == nil {
			log.Debugln("build type null while get default value", self.Name)
			return ""
		}

		if len(self.Complex.Fields) == 0 {
			return ""
		}

		return self.Complex.Fields[0].Name

	}

	return ""
}

func (self *FieldDescriptor) ListSpliter() string {

	return self.Meta.GetString("ListSpliter")
}

func (self *FieldDescriptor) RepeatCheck() bool {

	return self.Meta.GetBool("RepeatCheck")
}

var strByFieldDescriptor = map[FieldType]string{
	FieldType_None:   "none",
	FieldType_Int32:  "int32",
	FieldType_Int64:  "int64",
	FieldType_UInt32: "uint32",
	FieldType_UInt64: "uint64",

	FieldType_Float:  "float",
	FieldType_String: "string",
	FieldType_Bool:   "bool",
	FieldType_Enum:   "enum",
	FieldType_Struct: "struct",
	FieldType_Int:    "int",
	FieldType_Text:   "text",
	// FieldType_Vector3:    "Vector3",
	FieldType_Key:        "key",
	FieldType_CustomEnum: "customEnum",
	FieldType_Dict:       "dict",
	FieldType_Double:     "double",
}

var fieldTypeByString = make(map[string]FieldType)

func FieldTypeToString(t FieldType) string {
	if v, ok := strByFieldDescriptor[t]; ok {
		return v
	}

	return "unknown"
}

func ParseFieldType(str string) (t FieldType, ok bool) {
	v, ok := fieldTypeByString[str]
	return v, ok
}

const RepeatedKeyword = "repeated"
const RepeatedKeywordLen = len(RepeatedKeyword)

const SliceKeyword = "[]"
const SliceKeywordLen = len(SliceKeyword)

// 二维数组支持
const SliceExKeyWord = "[][]"
const SliceExKeyWordLen = len(SliceExKeyWord)

func (self *FieldDescriptor) ParseType(fileD *FileDescriptor, rawstr string) bool {
	var puretype string

	self.ExportClient = true
	self.ExportServer = true

	//首先 去掉类型中的|C 或者|S 并记录导出类型 分开导出
	if strings.Contains(rawstr, "|C") {
		self.ExportServer = false
		rawstr = strings.Replace(rawstr, "|C", "", 1)
	} else if strings.Contains(rawstr, "|S") {
		self.ExportClient = false
		rawstr = strings.Replace(rawstr, "|S", "", 1)
	}

	// 添加|key的支持
	if strings.Contains(rawstr, "|key") {
		self.IsKey = true
		rawstr = strings.Replace(rawstr, "|key", "", 1)
	} else if strings.Contains(rawstr, "key|") {
		self.IsKey = true
		rawstr = strings.Replace(rawstr, "key|", "", 1)
	} else if rawstr == "key" {
		self.IsKey = true
	} else {
		self.IsKey = false
	}

	if strings.HasPrefix(rawstr, RepeatedKeyword) {

		puretype = rawstr[RepeatedKeywordLen+1:]

		self.IsRepeated = true
	} else if strings.HasPrefix(rawstr, SliceExKeyWord) {
		// 解析二维数组类型 这里要先于一维数组
		puretype = rawstr[SliceExKeyWordLen:]
		self.Is2DArray = true
		self.IsRepeated = true
	} else if strings.HasSuffix(rawstr, SliceExKeyWord) {
		length := len(rawstr)
		puretype = rawstr[:length-SliceExKeyWordLen]

		self.Is2DArray = true
		self.IsRepeated = true
	} else if strings.HasPrefix(rawstr, SliceKeyword) {
		puretype = rawstr[SliceKeywordLen:]

		self.IsRepeated = true
	} else if strings.HasSuffix(rawstr, SliceKeyword) {
		length := len(rawstr)
		puretype = rawstr[:length-SliceKeywordLen]

		self.IsRepeated = true
	} else {
		puretype = rawstr
	}

	// 枚举类型 单独定义
	if strings.Contains(rawstr, "{") && strings.Contains(rawstr, "}") {
		self.EnumMap = make(map[string]int)
		valueList := strings.Split(rawstr, "\n")
		for i, v := range valueList {
			// 零值应该是枚举的名称 这里暂时不处理
			if i == 0 {
			} else {
				// 后面的元素应该是枚举值或末尾的大阔号
				v = strings.TrimSpace(v)
				v = strings.Replace(v, "{", "", 1)
				v = strings.Replace(v, "}", "", 1)
				if strings.Contains(v, "=") {
					keyValue := strings.Split(v, "=")
					if len(keyValue) != 2 {
						return false
					} else {
						key := strings.TrimSpace(keyValue[0])
						value := strings.TrimSpace(keyValue[1])
						nValue, _ := strconv.Atoi(value)
						self.EnumMap[key] = nValue
					}
				}
			}
		}
		self.IsCustomEnum = true
		puretype = "customEnum"
	}

	// 字典类型
	if strings.Index(rawstr, "[") == 0 && strings.Index(rawstr, "]") != 1 {
		rawstr = strings.TrimSpace(rawstr)
		rawstr = strings.Replace(rawstr, "[", "", 1)
		typeStrs := strings.Split(rawstr, "]")
		if len(typeStrs) == 2 {
			keyType, _ := ParseFieldType(typeStrs[0])
			valueType, _ := ParseFieldType(typeStrs[1])
			self.DictTypes[0] = keyType
			self.DictTypes[1] = valueType
			puretype = "dict"
		} else {
			log.Errorf("字典类型解析错误：%s", rawstr)
			return false
		}
	}

	// 这里得到类型的字符串
	if ft, ok := ParseFieldType(puretype); ok {
		self.Type = ft
		return true
	}

	if desc, ok := fileD.DescriptorByName[puretype]; ok {
		self.Complex = desc

		// 根据内建类型转成字段类型
		switch desc.Kind {
		case DescriptorKind_Struct:
			self.Type = FieldType_Struct
		case DescriptorKind_Enum:
			self.Type = FieldType_Enum
		}

	} else {
		return false
	}

	return true
}

func init() {

	for k, v := range strByFieldDescriptor {
		fieldTypeByString[v] = k
	}
}
