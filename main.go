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
	Version = "2.9.1"
)

func main() {

	flag.Parse()

	// 版本
	if *paramVersion {
		fmt.Printf("%s", Version)
		return
	}

	fileList := GetInputFileList(*paramPath)
	fmt.Println(fileList)

	for _, v := range fileList {
		switch *paramMode {
		case "v3":
			V3Entry()
		case "exportorv2", "v2":
			fileList := make([]string, 0)
			V2Entry(append(fileList, v))
		case "v2tov3":
			V2ToV3Entry()
		default:
			fmt.Println("--mode not specify")
			os.Exit(1)
		}
	}
}

func GetInputFileList(pathname string) []string {
	fileList := []string{}

	rd, _ := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		if fi.IsDir() {
			//fmt.Printf("[%s]\n", pathname+"\\"+fi.Name())
			list := GetInputFileList(pathname + "\\" + fi.Name())
			fileList = append(fileList, list...)
		} else {
			if !strings.Contains(fi.Name(), ".lua") {
				name := pathname + "\\" + fi.Name()
				fileList = append(fileList, name)
				fmt.Println(fileList)
			}
		}
	}

	return fileList
}
