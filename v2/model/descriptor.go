package model

import (
	"errors"
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

func (self *Descriptor) FieldByGlobalMap(value string, count int) *FieldDescriptor {
	mapList := self.FdMap[value]
	if mapList != nil {
		name := mapList[count]
		if name == "" {
			return nil
		}
		for _, v := range self.FieldByName {
			if v.Name == name {
				return v
			}

			if v.Meta.GetString("Alias") == name {
				return v
			}
		}
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
