package filter

import (
	"strconv"
	"strings"

	"github.com/davyxu/tabtoy/v2/i18n"
	"github.com/davyxu/tabtoy/v2/model"
)

var DictSpliter = "**++--**"

// 从单元格原始数据到最终输出的数值, 检查并转换, 处理默认值及根据meta转换情况
func ConvertValue(fd *model.FieldDescriptor, value string, fileD *model.FileDescriptor, node *model.Node) (ret string, ok bool) {
	// 空格, 且有默认值时, 使用默认值
	if value == "" {
		value = fd.DefaultValue()
	}

	switch fd.Type {
	case model.FieldType_Int32:
		_, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			log.Debugln(err)
			return "", false
		}

		ret = value
		node.AddValue(ret)
	// 添加新的支持类型 int默认为int64
	case model.FieldType_Int64, model.FieldType_Int:
		value = strings.TrimRight(value, "}")
		_, err := strconv.ParseInt(value, 10, 64)

		if err != nil {
			_, err := strconv.ParseFloat(value, 32)
			if err != nil {
				log.Debugln(err)
				return "", false
			}
		}

		ret = value
		node.AddValue(ret)
	case model.FieldType_UInt32:
		_, err := strconv.ParseUint(value, 10, 32)
		if err != nil {
			log.Debugln(err)
			return "", false
		}

		ret = value
		node.AddValue(ret)
	case model.FieldType_UInt64:
		_, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			log.Debugln(err)
			return "", false
		}

		ret = value
		node.AddValue(ret)
	case model.FieldType_Float, model.FieldType_Double:
		_, err := strconv.ParseFloat(value, 32)
		if err != nil {
			log.Debugln(err)
			return "", false
		}
		ret = value
		node.AddValue(ret)
	case model.FieldType_Bool:

		for {
			if value == "是" {
				ret = "true"
				break
			} else if value == "否" {
				ret = "false"
				break
			}

			v, err := strconv.ParseBool(value)

			if err != nil {
				log.Debugln(err)
				return "", false
			}

			if v {
				ret = "true"
			} else {
				ret = "false"
			}

			break
		}

		node.AddValue(ret)
	case model.FieldType_String:
		// 这里要检查中文字符
		if strings.Contains(value, "，") {
			log.Errorf("%s, '%s'", i18n.String(i18n.ConvertValue_CannotHaveCNComma), value)
			return "", false
		} else {
			ret = value
			node.AddValue(ret)
		}
	case model.FieldType_Text:
		ret = value
		node.AddValue(ret)
	case model.FieldType_Key:
		if value == "" {
			log.Errorf("%s, '%s'", i18n.String(i18n.ConvertValue_KeyNil), fd.Name)
			return "", false
		}

		// 暂时用int64的方式来解析key
		_, err := strconv.ParseInt(value, 10, 64)

		// 不等于0 则认为是str key
		if err != nil {
			ret = value
			node.AddValue(ret)
		}

		ret = value
		node.AddValue(ret)
	case model.FieldType_Enum:
		if fd.Complex == nil {
			log.Errorf("%s, '%s'", i18n.String(i18n.ConvertValue_EnumTypeNil), fd.Name)
			return "", false
		}

		evd := fd.Complex.FieldByValueAndMeta(value)
		if evd == nil {
			log.Errorf("%s, '%s' '%s'", i18n.String(i18n.ConvertValue_EnumValueNotFound), value, fd.Complex.Name)
			return "", false
		}

		// 使用枚举的英文字段名输出
		ret = evd.Name
		node.AddValue(ret).EnumValue = evd.EnumValue

	case model.FieldType_Struct:
		if fd.Complex == nil {
			log.Errorf("%s, '%s'", i18n.String(i18n.ConvertValue_StructTypeNil), fd.Name)
			return "", false
		}

		if value == "" {
			node.IsEmpty = true
			if !fillStructDefaultValue(fd.Complex, fileD, node) {
				return "", false
			}

		} else {
			if !parseStruct(fd, value, fileD, node) {
				return "", false
			}
		}
	case model.FieldType_CustomEnum:
		value = strings.TrimSpace(value)
		mapValue := fd.EnumMap[value]
		// 如果没找到 就直接导出原值
		if mapValue == 0 {
			ret = value
		} else {
			ret = strconv.Itoa(fd.EnumMap[value])
		}
		node.AddValue(ret)
	case model.FieldType_Dict:
		value = strings.TrimSpace(value)
		value = strings.Replace(value, "{", "", 1)
		value = strings.Replace(value, "}", "", 1)
		ret = value
		node.AddValue(ret)
	default:
		log.Errorf("%s, '%s' '%s'", i18n.String(i18n.ConvertValue_UnknownFieldType), fd.Name, fd.Name)
		return "", false
	}

	ok = true

	return
}

// func convertValueForDict(fType model.FieldType, value string) string {
// 	switch fType {
// 	case model.FieldType_Int32:
// 		_, err := strconv.ParseInt(value, 10, 32)
// 		if err != nil {
// 			log.Debugln(err)
// 			return ""
// 		}

// 		return value
// 	// 添加新的支持类型 int默认为int64
// 	case model.FieldType_Int64, model.FieldType_Int:
// 		value = strings.TrimRight(value, "}")
// 		_, err := strconv.ParseInt(value, 10, 64)

// 		if err != nil {
// 			_, err := strconv.ParseFloat(value, 32)
// 			if err != nil {
// 				log.Debugln(err)
// 				return ""
// 			}
// 		}

// 		return value
// 	case model.FieldType_UInt32:
// 		_, err := strconv.ParseUint(value, 10, 32)
// 		if err != nil {
// 			log.Debugln(err)
// 			return ""
// 		}

// 		return value
// 	case model.FieldType_UInt64:
// 		_, err := strconv.ParseUint(value, 10, 64)
// 		if err != nil {
// 			log.Debugln(err)
// 			return ""
// 		}

// 		return value
// 	case model.FieldType_Float:
// 		_, err := strconv.ParseFloat(value, 32)
// 		if err != nil {
// 			log.Debugln(err)
// 			return ""
// 		}
// 		return value
// 	case model.FieldType_Bool:
// 		var ret string
// 		for {
// 			if value == "是" {
// 				ret = "true"
// 				break
// 			} else if value == "否" {
// 				ret = "false"
// 				break
// 			}

// 			v, err := strconv.ParseBool(value)

// 			if err != nil {
// 				log.Debugln(err)
// 				return ""
// 			}

// 			if v {
// 				ret = "true"
// 			} else {
// 				ret = "false"
// 			}
// 			break
// 		}

// 		return ret
// 	case model.FieldType_String:
// 		// 这里要检查中文字符
// 		if strings.Contains(value, "，") {
// 			log.Errorf("%s, '%s'", i18n.String(i18n.ConvertValue_CannotHaveCNComma), value)
// 			return ""
// 		} else {
// 			return value
// 		}
// 	case model.FieldType_Text:
// 		return value
// 	default:
// 		return ""
// 	}
// }

func parseVectors(value string) []string {
	value = strings.Trim(value, "{")
	value = strings.Trim(value, "}")
	value = strings.Replace(value, " ", "", -1)
	value = strings.Replace(value, "x=", "", 1)
	value = strings.Replace(value, "y=", "", 1)
	value = strings.Replace(value, "z=", "", 1)
	values := strings.Split(value, ",")

	return values
}

// 填充空结构体的默认值
func fillStructDefaultValue(structD *model.Descriptor, fileD *model.FileDescriptor, node *model.Node) bool {

	for _, fd := range structD.Fields {

		// 没默认值不输出, 建议忽略的字段除外, 先导出node, 再在printer中忽略
		if fd.Meta.GetString("Default") == "" && node.SugguestIgnore {
			continue
		}

		fieldNode := node.AddKey(fd)

		// 结构体的值没填, 且没默认值, 建议忽略
		if fd.Meta.GetString("Default") == "" && node.Value == "" {
			fieldNode.SugguestIgnore = true
		}

		_, ok := ConvertValue(fd, "", fileD, fieldNode)
		if !ok {
			return false
		}
	}

	return true

}
