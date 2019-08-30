package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/davyxu/golog"
)

var log = golog.New("main")

const (
	Version        = "2.9.1"
	MutileFileFlag = "_"
)

func main() {

	flag.Parse()

	// 版本
	if *paramVersion {
		fmt.Printf("%s", Version)
		return
	}

	fileList := GetInputFileList(*paramPath)

	for _, v := range fileList {
		// 无下划线 则单文件直接导出
		files := []string{}
		RunProgram(append(files, v))
	}
}

// 递归读取表格
func GetInputFileList(pathname string) []string {
	fileList := []string{}

	rd, _ := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		if fi.IsDir() {
			//fmt.Printf("[%s]\n", pathname+"\\"+fi.Name())
			list := GetInputFileList(pathname + "\\" + fi.Name())
			fileList = append(fileList, list...)
		} else {
			// 忽略掉脚本文件
			if !strings.Contains(fi.Name(), ".lua") {
				name := pathname + "\\" + fi.Name()
				// 如果多文件的话 则尝试在列表中找出
				if strings.Contains(name, MutileFileFlag) {
					fileList = GetCombinedFileList(fileList, name)
				} else {
					fileList = append(fileList, name)
				}
			}
		}
	}

	return fileList
}

func RunProgram(fileList []string) {
	switch *paramMode {
	case "v3":
		V3Entry()
	case "exportorv2", "v2":
		V2Entry(fileList)
	case "v2tov3":
		V2ToV3Entry()
	default:
		fmt.Println("--mode not specify")
		os.Exit(1)
	}
}

func GetCombinedFileList(fileList []string, newFile string) []string {
	// 首先对名称进行切割，切掉带有下划线的部分
	name := strings.Split(newFile, MutileFileFlag)[0]

	for i, v := range fileList {
		// 主文件是不包括"_"的
		if !strings.Contains(v, MutileFileFlag) {
			vName := strings.Replace(v, ".xlsx", "", 1)
			vName = strings.Replace(vName, ".csv", "", 1)
			if vName == name {
				fileList[i] = fileList[i] + "+" + newFile
				return fileList
			}
		}
	}
	// 没有找到主文件 直接加到列表中
	return append(fileList, newFile)
}
