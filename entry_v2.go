package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	v2 "github.com/davyxu/tabtoy/v2"
	"github.com/davyxu/tabtoy/v2/i18n"
	"github.com/davyxu/tabtoy/v2/printer"
)

// v2特有
var (
	paramProtoVersion = flag.Int("protover", 3, "output .proto file version, 2 or 3")

	paramLuaEnumIntValue = flag.Bool("luaenumintvalue", true, "use int type in lua enum value")
	paramLuaTabHeader    = flag.String("luatabheader", "", "output string to lua tab header")

	paramGenCSharpBinarySerializeCode = flag.Bool("cs_gensercode", true, "generate c# binary serialize code, default is true")
)

func V2Entry() {
	g := printer.NewGlobals()

	if *paramLanguage != "" {
		if !i18n.SetLanguage(*paramLanguage) {
			log.Infof("language not support: %s", *paramLanguage)
		}
	}

	g.Version = Version

	// 可自动加载Global类型表 这里单独循环是必要的的 因为Global要加到前面才行
	hasGlobal := false
	for _, v := range flag.Args() {
		if v == "Globals.xlsx" {
			hasGlobal = true
		}
	}

	if !hasGlobal {
		g.InputFileList = append(g.InputFileList, "Globals.xlsx")
	}

	for _, v := range flag.Args() {
		g.InputFileList = append(g.InputFileList, v)
	}

	g.ParaMode = *paramPara
	g.CombineStructName = *paramCombineStructName
	g.ProtoVersion = *paramProtoVersion
	g.LuaEnumIntValue = *paramLuaEnumIntValue
	g.LuaTabHeader = *paramLuaTabHeader
	g.GenCSSerailizeCode = *paramGenCSharpBinarySerializeCode
	g.PackageName = *paramPackageName
	g.Path = *paramPath

	//fileList := GetInputFileList(g.Path)
	if *paramProtoOut != "" {
		g.AddOutputType("proto", *paramProtoOut)
	}

	if *paramPbtOut != "" {
		g.AddOutputType("pbt", *paramPbtOut)
	}

	if *paramJsonOut != "" {
		g.AddOutputType("json", *paramJsonOut)
	}

	fileList := make([]string, 0)
	for _, v := range g.InputFileList {
		fileList = append(fileList, v.(string))
	}

	if *paramLuaOut != "" {
		g.AddOutputType("lua", ParseFileList(fileList))
	}

	if *paramCSharpOut != "" {
		g.AddOutputType("cs", *paramCSharpOut)
	}

	if *paramGoOut != "" {
		g.AddOutputType("go", *paramGoOut)
	}

	if *paramCppOut != "" {
		g.AddOutputType("cpp", *paramCppOut)
	}

	if *paramBinaryOut != "" {
		g.AddOutputType("bin", *paramBinaryOut)
	}

	if *paramTypeOut != "" {
		g.AddOutputType("type", *paramTypeOut)
	}

	if !v2.Run(g) {
		os.Exit(1)
	}
}

func GetInputFileList(pathname string) []string {
	fileList := make([]string, 0)
	rd, _ := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		if fi.IsDir() {
			// fmt.Printf("[%s]\n", pathname+"\\"+fi.Name())
			GetInputFileList(pathname + "\\" + fi.Name())
		} else {
			fmt.Println(fi.Name())

			if !strings.Contains(fi.Name(), ".lua") {
				fileList = append(fileList, pathname+"\\"+fi.Name())
			}
		}
	}
	return fileList
}

func ParseFileList(fileList []string) string {
	for _, v := range fileList {
		if v != "Globals.xlsx" {
			fmt.Println("v:" + v)
			name := strings.Replace(v, ".xlsx", "", 1)
			name = strings.Replace(name, ".csv", "", 1)
			return name + ".lua"
		}
	}
	return ""
}
