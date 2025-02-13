package cons

import (
	"fmt"
	"strings"
)

//环境引用
// true 静态文件
// false 打包二进制文佳 (要求打包html目录)
var isStatic = true

//初始化 扫描路径
var BaseUrl = "https://www.cdnbus.one/"
var DirFile = ""
var BaseDir = map[string]string{
	//"1":  "E:\\emby",
	//"2":  "e:\\emby\\tomake",
	//"3":  "F:\\emby\\tomake",
	//"4": "F:\\emby\\emby-rename",
	//"5":  "H:\\emby\\tomake",
	//"6":  "H:\\emby\\emby-rename",
	//"7": "I:\\emby\\tomake",
	//"8": "I:\\emby\\emby-rename",
	//"9":  "g:\\emby\\tomake",
	//"10": "g:\\emby\\emby-rename",
}
var Images = []string{}
var Docs = []string{}
var VideoTypes = []string{}
var QueryTypes = []string{}

func SetBaseDir(dirs []string) {
	BaseDir = make(map[string]string)
	for _, name := range dirs {
		index := strings.ReplaceAll(name, "\\", "~")
		fmt.Println(index, ":", name)
		BaseDir[string(index)] = name
	}
}
