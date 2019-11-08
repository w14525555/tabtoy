package v2

import (
	"strings"

	"github.com/davyxu/tabtoy/v2/filter"
	"github.com/davyxu/tabtoy/v2/i18n"
	"github.com/davyxu/tabtoy/v2/model"
)

func coloumnProcessor(file model.GlobalChecker, record *model.Record, fd *model.FieldDescriptor, raw string, sugguestIgnore bool) bool {

	spliter := fd.ListSpliter()

	if fd.IsRepeated {
		if spliter == "" {
			spliter = ","
		}

		var valueList []string
		raw = strings.ReplaceAll(raw, " ", "")
		if fd.Type == model.FieldType_Struct || fd.Is2DArray {
			// 结构体数组分割
			raw = strings.TrimSpace(raw)
			raw = strings.Replace(raw, "},{", "}{", -1)
			valueList = strings.Split(raw, "}{")
		} else {
			valueList = parseVector(fd.Type, raw)
		}

		var node *model.Node

		if fd.Type != model.FieldType_Struct {
			node = record.NewNodeByDefine(fd)
		}

		for _, v := range valueList {

			rawSingle := strings.TrimSpace(v)

			// 结构体要多添加一个节点, 处理repeated 结构体情况
			if fd.Type == model.FieldType_Struct {
				node = record.NewNodeByDefine(fd)
				node.StructRoot = true
				node = node.AddKey(fd)
			}

			if raw != "" {
				// 二维数组仍旧需要再次分割
				if fd.Is2DArray {
					var subValueList []string
					rawNew := strings.TrimLeft(rawSingle, "{")
					rawNew = strings.TrimRight(rawNew, "}")
					// 如果最后一位是逗号 则需要干掉这个逗号 否则分割会错误
					rawNew = strings.TrimRight(rawNew, "},")

					//fmt.Println(rawNew)
					subValueList = strings.Split(rawNew, spliter)
					newNode := record.NewNodeByDefine(fd)
					newNode.AddValue("{")
					for _, vv := range subValueList {
						dataProcessor(file, fd, vv, node)
					}
					newNode = record.NewNodeByDefine(fd)
					newNode.AddValue("}")
				} else if !dataProcessor(file, fd, rawSingle, node) {
					return false
				}
			}

		}

	} else { // 普通数据/repeated单元格分多个列

		node := record.NewNodeByDefine(fd)

		node.SugguestIgnore = sugguestIgnore

		// 结构体要多添加一个节点, 处理repeated 结构体情况
		if fd.Type == model.FieldType_Struct {
			node.StructRoot = true
			node = node.AddKey(fd)
		}

		node.SugguestIgnore = sugguestIgnore

		if !dataProcessor(file, fd, raw, node) {
			return false
		}
	}

	return true
}

// 解析vector类型
func parseVector(ft model.FieldType, raw string) []string {
	raw = strings.Replace(raw, "{", "", -1)
	raw = strings.Replace(raw, "}", "", -1)
	valueList := strings.Split(raw, ",")

	// if ft == model.FieldType_Vector2 {
	// 	return getVectorResult(2, valueList)
	// } else if ft == model.FieldType_Vector3 {
	// 	return getVectorResult(3, valueList)
	// }

	return valueList
}

func getVectorResult(vectorNum int, valueList []string) []string {
	//  valueList[0] != "" 对应空的字符串数组的情况 空字符串数组长度为1
	if len(valueList)%vectorNum != 0 && valueList[0] != "" {
		log.Errorf("%s, '%s'", i18n.String(i18n.ConvertValue_VectorError), valueList)
	}

	var result []string
	for i, v := range valueList {
		if i%vectorNum == 0 {
			result = append(result, v)
		} else {
			result[i/vectorNum] = result[i/vectorNum] + "," + v
		}
	}
	return result
}

func dataProcessor(gc model.GlobalChecker, fd *model.FieldDescriptor, raw string, node *model.Node) bool {
	// 单值
	if cv, ok := filter.ConvertValue(fd, raw, gc.GlobalFileDesc(), node); !ok {
		goto ConvertError

	} else {

		// 值重复检查
		if (fd.Meta.GetBool("RepeatCheck") || fd.Name == "key" || fd.IsKey) && !gc.CheckValueRepeat(fd, cv) {
			log.Errorf("%s, %s raw: '%s'", i18n.String(i18n.DataSheet_ValueRepeated), fd.String(), cv)
			return false
		}
	}

	return true

ConvertError:

	log.Errorf("%s, %s raw: '%s'", i18n.String(i18n.DataSheet_ValueConvertError), fd.String(), raw)

	return false
}
