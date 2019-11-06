package v2

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"github.com/davyxu/tabtoy/v2/i18n"
	"github.com/davyxu/tabtoy/v2/model"
	"github.com/davyxu/tabtoy/v2/printer"
	"github.com/tealeg/xlsx"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// 检查单元格值重复结构
type valueRepeatData struct {
	fd    *model.FieldDescriptor
	value string
}

// 1个电子表格文件
type File struct {
	LocalFD  *model.FileDescriptor // 本文件的类型描述表
	GlobalFD *model.FileDescriptor // 全局的类型描述表
	FileName string
	coreFile *xlsx.File

	dataSheets  []*DataSheet
	Header      *DataHeader
	dataHeaders []*DataHeader

	valueRepByKey map[valueRepeatData]bool // 检查单元格值重复map

	mergeList []*File
}

func (self *File) GlobalFileDesc() *model.FileDescriptor {
	return self.GlobalFD

}

// 解析表格 判定是类型表还是数据表
func (self *File) ExportLocalType(mainFile *File, g *printer.Globals, fileList []string) bool {

	var sheetCount int

	var typeSheet *TypeSheet
	// 解析类型表
	for _, rawSheet := range self.coreFile.Sheets {

		if isTypeSheet(rawSheet.Name) {
			if sheetCount > 0 {
				log.Errorf("%s", i18n.String(i18n.File_TypeSheetKeepSingleton))
				return false
			}
			typeSheet = newTypeSheet(NewSheet(self, rawSheet))

			// 从cell添加类型
			if !typeSheet.Parse(self.LocalFD, self.GlobalFD) {
				return false
			}

			sheetCount++
		}
	}

	// if typeSheet == nil {
	// 	log.Errorf("%s", i18n.String(i18n.File_TypeSheetNotFound))
	// 	return false
	// }

	// 解析表头
	for _, rawSheet := range self.coreFile.Sheets {

		// 是数据表
		if !isTypeSheet(rawSheet.Name) {
			dSheet := newDataSheet(NewSheet(self, rawSheet))

			if !dSheet.Valid() {
				continue
			}

			log.Infof("%s", rawSheet.Name)

			dataHeader := newDataHeadSheet()

			// 检查引导头
			if !dataHeader.ParseProtoField(len(self.dataSheets), dSheet.Sheet, self.LocalFD, self.GlobalFD, g, fileList) {
				return false
			}

			if mainFile != nil {

				if fieldName, ok := dataHeader.AsymmetricEqual(mainFile.Header); !ok {
					log.Errorf("%s main: %s child: %s field: %s", i18n.String(i18n.DataHeader_NotMatchInMultiTableMode), mainFile.FileName, self.FileName, fieldName)
					return false
				}

			}

			if self.Header == nil {
				self.Header = dataHeader
			}

			self.dataHeaders = append(self.dataHeaders, dataHeader)
			self.dataSheets = append(self.dataSheets, dSheet)

		}
	}

	// File描述符的名字必须放在类型里, 因为这里始终会被调用, 但是如果数据表缺失, 是不会更新Name的
	self.LocalFD.Name = self.LocalFD.Pragma.GetString("TableName")

	return true
}

func (self *File) IsVertical() bool {
	return self.LocalFD.Pragma.GetBool("Vertical")
}

func (self *File) ExportData(dataModel *model.DataModel, parentHeader *DataHeader) bool {

	for index, d := range self.dataSheets {

		//log.Infof("%s", d.Name)

		// 多个sheet时, 使用和多文件一样的父级
		if parentHeader == nil && len(self.dataHeaders) > 1 {
			parentHeader = self.dataHeaders[0]
		}

		if !d.Export(self, dataModel, self.dataHeaders[index], parentHeader) {
			return false
		}
	}

	return true

}

func (self *File) CheckValueRepeat(fd *model.FieldDescriptor, value string) bool {

	key := valueRepeatData{
		fd:    fd,
		value: value,
	}

	if _, ok := self.valueRepByKey[key]; ok {
		return false
	}

	self.valueRepByKey[key] = true

	return true
}

func isTypeSheet(name string) bool {
	return strings.TrimSpace(name) == model.TypeSheetName
}

func NewFile(filename string) *File {

	self := &File{
		valueRepByKey: make(map[valueRepeatData]bool),
		LocalFD:       model.NewFileDescriptor(),
		FileName:      filename,
	}

	var err error
	if strings.HasSuffix(filename, "xlsx") {
		self.coreFile, err = xlsx.OpenFile(filename)
	} else {
		self.coreFile = generateXLSXFromCSV(filename, ",")
		if self.coreFile == nil {
			fmt.Println("NAnui")
		}
		if err != nil {
			fmt.Printf(err.Error())
			return nil
		}
	}

	if err != nil {
		log.Errorln(err.Error())
		log.Errorf("%s, %v", i18n.String(i18n.System_OpenReadXlsxFailed), err.Error())

		return nil
	}

	return self
}

func generateXLSXFromCSV(csvPath string, delimiter string) *xlsx.File {
	csvFile, err := os.Open(csvPath)
	if err != nil {
		fmt.Printf(err.Error())
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	if len(delimiter) > 0 {
		reader.Comma = rune(delimiter[0])
	} else {
		reader.Comma = rune(',')
	}
	xl := xlsx.NewFile()
	_, name := filepath.Split(csvPath)
	sheet, err := xl.AddSheet(name)
	if err != nil {
		fmt.Printf(err.Error())
	}
	fields, err := reader.Read()

	for err == nil {
		row := sheet.AddRow()
		for _, field := range fields {
			field = strings.TrimFunc(field, IsBom)
			// 如果不符合UTF8标准 则进行转换 注意只能支持从GBK格式的转换
			if len(field) > 0 && !utf8.ValidString(field) {
				data, _ := ConvGBKToUTF8([]byte(field))
				field = string(data)
			}
			cell := row.AddCell()
			cell.Value = field
		}
		fields, err = reader.Read()
	}
	if err != nil {
		fmt.Printf(err.Error())
	}
	return xl
}

func ConvGBKToUTF8(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func ConvUTF8ToGBK(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func IsBom(r rune) bool {
	if uint32(r) == 65279 {
		return true
	} else {
		return false
	}
}

func ParseCSVName() {

}
