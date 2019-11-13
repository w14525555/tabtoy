package v2

import (
	"fmt"
	"strings"

	"github.com/davyxu/tabtoy/util"
	"github.com/davyxu/tabtoy/v2/i18n"
	"github.com/davyxu/tabtoy/v2/model"
	"github.com/davyxu/tabtoy/v2/printer"
)

/*
	Sheet数据表单类型头

*/

const (
	// 信息所在的行
	DataSheetHeader_FieldName = 3 // 字段名(对应proto)
	DataSheetHeader_FieldType = 2 // 字段类型
	DataSheetHeader_FieldMeta = 0 // 字段特性
	DataSheetHeader_Comment   = 1 // 用户注释
	DataSheetHeader_DataBegin = 4 // 数据开始
)

type DataHeader struct {

	// 按排列的, 保留有注释掉的字段和重复的repeated列
	rawHeaderFields []*model.FieldDescriptor

	// 按排列的, 字段不重复
	headerFields []*model.FieldDescriptor

	// 按字段名分组索引字段, 字段不重复
	HeaderByName map[string]*model.FieldDescriptor
}

// 检查字段行的长度
func (self *DataHeader) ParseProtoField(index int, sheet *Sheet, localFD *model.FileDescriptor, globalFD *model.FileDescriptor, g *printer.Globals, fileList []string) bool {

	verticalHeader := localFD.Pragma.GetBool("Vertical")

	// fmt.Println(globalFD.Name)

	// 适用于配置的表格
	if verticalHeader {
		// 遍历行(从第二行开始)
		for sheet.Row = 1; ; sheet.Row++ {

			he := &DataHeaderElement{
				FieldName: sheet.GetCellData(sheet.Row, DataSheetHeader_FieldName),
				FieldType: sheet.GetCellData(sheet.Row, DataSheetHeader_FieldType),
				FieldMeta: sheet.GetCellData(sheet.Row, DataSheetHeader_FieldMeta),
				Comment:   sheet.GetCellData(sheet.Row, DataSheetHeader_Comment),
			}

			if he.FieldName == "" {
				break
			}

			if errorPos := self.addHeaderElement(he, localFD, globalFD); errorPos != -1 {
				sheet.Column = errorPos
				goto ErrorStop
			}

		}

	} else {
		// 适用于正常数据的表格

		// 遍历列
		for sheet.Column = 0; ; sheet.Column++ {

			he := &DataHeaderElement{
				FieldName: sheet.GetCellData(DataSheetHeader_FieldName, sheet.Column),
				FieldType: sheet.GetCellData(DataSheetHeader_FieldType, sheet.Column),
				FieldMeta: sheet.GetCellData(DataSheetHeader_FieldMeta, sheet.Column),
				Comment:   sheet.GetCellData(DataSheetHeader_Comment, sheet.Column),
			}

			// 第一列第一行不支持meta 因为要存放介绍等信息
			// 以后可能支持 但是由于比较复杂 暂且不做
			if sheet.Column == 0 {
				meta := he.FieldMeta
				he.FieldMeta = ""
				fmt.Println("Name: " + sheet.file.FileName)
				if !CheckOutputType(g, meta, fileList) {
					return false
				}
			}

			if he.FieldName == "" && he.Comment == "" {
				break
			}

			// 默认类型为int
			if he.FieldType == "" {
				he.FieldType = "int"
			}

			if errorPos := self.addHeaderElement(he, localFD, globalFD); errorPos != -1 {
				sheet.Row = errorPos
				goto ErrorStop
			}

		}

	}

	if len(self.rawHeaderFields) == 0 {
		return false
	}

	if index == 0 {
		// 添加第一个数据表的定义
		if !self.makeRowDescriptor(localFD, self.headerFields, globalFD.Fdmap) {
			goto ErrorStop
		}
	}

	return true

ErrorStop:

	r, c := sheet.GetRC()

	log.Errorf("%s|%s(%s)", sheet.file.FileName, sheet.Name, util.R1C1ToA1(r, c))
	return false
}

func CheckOutputType(g *printer.Globals, meta string, fileList []string) bool {
	// 先判定是否已经读取文件的导出类型
	if !g.HasReadExportType {
		// 这里要记录文件导出的模式
		metas := strings.Split(meta, " ")
		// 客户端导出
		client := true
		server := true

		// 逗号分隔的
		metaSplitByComma := make([]string, 0, 0)
		// 因为可能存在着逗号分割的情况 因此要兼容
		for _, v := range metas {
			if strings.Contains(v, ",") {
				values := strings.Split(v, ",")
				metaSplitByComma = append(metaSplitByComma, values...)
			} else {
				metaSplitByComma = append(metaSplitByComma, v)
			}
		}

		// 获取客户端和服务器是否导出
		for _, v := range metaSplitByComma {
			if strings.Contains(v, "C=") {
				index := strings.Index(v, "C=")
				var value string
				if index != 0 {
					// 如果C=前面还有别的字符 那就截取掉
					value = v[index:]
				} else {
					value = v
				}
				value = strings.Replace(value, "C=", "", 1)
				if value == "lua" {
					client = true
				} else {
					client = false
				}
			}

			if strings.Contains(v, "S=") {
				index := strings.Index(v, "S=")
				var value string
				if index != 0 {
					// 如果C=前面还有别的字符 那就截取掉
					value = v[index:]
				} else {
					value = v
				}

				value = strings.Replace(value, "S=", "", 1)
				if value == "lua" {
					server = true
				} else {
					server = false
				}
				// } else {
				// 	log.Errorf("不支持的导出类型：%s", value)
				// 	return false
				// }
			}
		}

		fileName := ParseFileList(fileList)
		// 根据是否导出增加OutPut
		if client {
			g.AddOutputType("lua", GetClientPath(fileName, g), 1)
		}

		if server {
			g.AddOutputType("lua", GetServerPath(fileName, g), 2)
		}

		g.HasReadExportType = true
	}

	return true
}

func (self *DataHeader) RawField(index int) *model.FieldDescriptor {
	if index >= len(self.rawHeaderFields) {
		return nil
	}

	return self.rawHeaderFields[index]
}

func (self *DataHeader) RawFieldCount() int {
	return len(self.rawHeaderFields)
}

func (self *DataHeader) FieldRepeatedCount(fd *model.FieldDescriptor) (count int) {

	for _, libfd := range self.rawHeaderFields {
		if libfd == fd {
			count++
		}
	}

	return

}

func (self *DataHeader) Equal(other *DataHeader) (string, bool) {

	if len(self.headerFields) != len(other.headerFields) {
		return "field len", false
	}

	for k, v := range self.headerFields {
		if !v.Equal(other.headerFields[k]) {
			return v.Name, false
		}
	}

	return "", true
}

func (self *DataHeader) AsymmetricEqual(other *DataHeader) (string, bool) {

	for _, otherFD := range other.headerFields {

		if otherFD == nil {
			continue
		}

		if thisFD, ok := self.HeaderByName[otherFD.Name]; ok {

			if !thisFD.Equal(otherFD) {
				return otherFD.Name, false
			}
		}

	}

	return "", true
}

func (self *DataHeader) addHeaderElement(he *DataHeaderElement, localFD *model.FileDescriptor, globalFD *model.FileDescriptor) int {
	def := model.NewFieldDescriptor()
	def.Name = he.FieldName

	var errorPos int = -1

	// #开头表示注释, 跳过
	if strings.Index(he.FieldName, "#") != 0 && he.FieldName != "" {

		errorPos = he.Parse(def, localFD, globalFD, self.HeaderByName)
		if errorPos != -1 {
			return errorPos
		}

		// 根据字段名查找, 处理repeated字段case
		exist, ok := self.HeaderByName[def.Name]

		if ok {

			errorPos = checkSameNameElement(exist, def)
			if errorPos != -1 {
				return errorPos
			}

			def = exist

		} else {

			// 普通表头合法性检查
			errorPos = checkElement(def)
			if errorPos != -1 {
				return errorPos
			}

			self.HeaderByName[def.Name] = def
			self.headerFields = append(self.headerFields, def)
		}
	}

	// 有注释字段, 但是依然要放到这里来进行索引
	self.rawHeaderFields = append(self.rawHeaderFields, def)

	return -1
}

func (self *DataHeader) makeRowDescriptor(fileD *model.FileDescriptor, rootField []*model.FieldDescriptor, fdmap map[string][]string) bool {

	rowType := model.NewDescriptor()
	rowType.Usage = model.DescriptorUsage_RowType
	rowType.Name = fmt.Sprintf("%sDefine", fileD.Pragma.GetString("TableName"))
	rowType.Kind = model.DescriptorKind_Struct

	// 类型已经存在, 说明是自己定义的 XXDefine, 不允许
	if _, ok := fileD.DescriptorByName[rowType.Name]; ok {
		log.Errorf("%s '%s'", i18n.String(i18n.DataHeader_UseReservedTypeName), rowType.Name)
		return false
	}

	fileD.Add(rowType)

	// 将表格中的列添加到类型中, 方便导出
	for _, field := range rootField {

		rowType.Add(field, fdmap)
	}

	return true

}

func newDataHeadSheet() *DataHeader {

	return &DataHeader{
		HeaderByName: make(map[string]*model.FieldDescriptor),
	}
}

// 解析得到导出文件的名称
func ParseFileList(fileList []string) string {
	for _, v := range fileList {
		if v != "Globals.xlsx" {
			name := strings.Replace(v, ".xlsx", "", 1)
			name = strings.Replace(name, ".csv", "", 1)
			name = strings.Split(name, "+")[0]
			return name + ".lua"
		}
	}
	return ""
}

func GetServerPath(name string, g *printer.Globals) string {
	return strings.Replace(name, g.Path, g.ServerOut, 1)
}

func GetClientPath(name string, g *printer.Globals) string {
	return strings.Replace(name, g.Path, g.ClientOut, 1)
}
