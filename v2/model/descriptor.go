package model

import (
	"errors"
	"strings"
)

type DescriptorKind int

const (
	DescriptorKind_None DescriptorKind = iota
	DescriptorKind_Enum
	DescriptorKind_Struct
)

type DescriptorUsage int

const (
	DescriptorUsage_None          DescriptorUsage = iota
	DescriptorUsage_RowType                       // 每个表的行类型
	DescriptorUsage_CombineStruct                 // 最终使用的合并结构体
)

type Descriptor struct {
	Name  string
	Kind  DescriptorKind
	Usage DescriptorUsage

	FieldByName   map[string]*FieldDescriptor
	FieldByIndex  map[int32]*FieldDescriptor
	FieldByNumber map[int32]*FieldDescriptor
	Fields        []*FieldDescriptor

	Indexes     []*FieldDescriptor
	IndexByName map[string]*FieldDescriptor

	File  *FileDescriptor
	FdMap map[string][]string
}

var (
	ErrDuplicateFieldName = errors.New("Duplicate field name")
	ErrDuplicateIndexName = errors.New("Duplicate index name")
)

func (self *Descriptor) Add(def *FieldDescriptor, fdmap map[string][]string) error {

	def.Parent = self
	def.Order = int32(len(self.Fields))

	// 创建字段
	if _, ok := self.FieldByName[def.Name]; ok {
		return ErrDuplicateFieldName
	} else {
		self.FieldByName[def.Name] = def
		self.FieldByNumber[def.EnumValue] = def
		self.Fields = append(self.Fields, def)
	}

	// 创建索引
	if def.Meta.GetBool("MakeIndex") {

		if _, ok := self.IndexByName[def.Name]; ok {
			return ErrDuplicateIndexName
		} else {
			self.IndexByName[def.Name] = def
			self.Indexes = append(self.Indexes, def)
		}
	}

	self.FdMap = fdmap

	return nil
}

func (self *Descriptor) FieldByValueAndMeta(value string) *FieldDescriptor {
	for _, v := range self.FieldByName {
		if v.Name == value {
			return v
		}

		if v.Meta.GetString("Alias") == value {
			return v
		}

	}

	return nil
}

// 结构体寻找
func (self *Descriptor) FieldByGlobalMap(key string, value string, count int, subFieldName string) *FieldDescriptor {
	mapList := self.FdMap[key]
	// fmt.Println(key, value, count, subFieldName)
	if mapList != nil {
		// 根据是第几个元素去读取字段名称
		var name string

		subFieldName = strings.TrimSpace(subFieldName)
		if len(subFieldName) > 0 {
			name = subFieldName
		} else {
			name = mapList[count]
			if name == "" {
				return nil
			}
		}

		var sample FieldDescriptor
		count := 0
		// fmt.Println(key, self.FieldByName)
		for _, v := range self.FieldByName {
			if v.Name == name {
				return v
			}

			if v.Meta.GetString("Alias") == name {
				return v
			}

			if count == 0 {
				sample = *v
			}
			count++
		}
		copy := NewFieldDescriptor()
		// fmt.Println(subFieldName, sample.Type)
		copy.Name = name
		copy.Type = sample.Type
		// fmt.Println(copy)
		return copy
	}

	return nil
}

func NewDescriptor() *Descriptor {
	return &Descriptor{
		FieldByName:   make(map[string]*FieldDescriptor),
		FieldByNumber: make(map[int32]*FieldDescriptor),
		IndexByName:   make(map[string]*FieldDescriptor),
	}
}
