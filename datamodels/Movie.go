package datamodels

import (
	"fmt"
	"time"

	"../utils"
)

//声明一个File结构体 表示一个文件信息
type Movie struct {
	Id       string
	Code     string "json:'code'"
	Name     string
	Path     string
	Png      string
	Nfo      string
	Jpg      string
	Actress  string
	FileType string
	DirPath  string
	Size     int64
	SizeStr  string
	CTime    string
	MTime    string
	PTime    string
	Studio   string
	Supplier string
	Length   string
	Series   string
	Director string
	Title    string
}

func NewFile(dir string, path string, name string, fileType string, size int64, modTime time.Time) Movie {
	// 使用工厂模式 返回一个 Movie 实例
	id, _ := utils.DirpathForId(path)
	result := Movie{
		Id:       id,
		Code:     utils.GetCode(name),
		Name:     name,
		Path:     path,
		Png:      utils.GetPng(path, "png"),
		Nfo:      utils.GetPng(path, "nfo"),
		Jpg:      utils.GetPng(path, "jpg"),
		Actress:  utils.GetActress(name),
		FileType: fileType,
		DirPath:  dir,
		Size:     size,
		SizeStr:  utils.GetSizeStr(size),
		CTime:    "",
		MTime:    modTime.Format("2006-01-02 15:04:05"),
	}
	return result
}

func (f Movie) GetFileInfo() string {
	//
	info := fmt.Sprintf("name: %v\t code:%v\t fileType:%v\t sizeStr:%v\t actress:%v\t path:%v\t",
		f.Name, f.Code, f.FileType, f.SizeStr, f.Actress, f.Path)
	return info
}
func (f Movie) PngBase64() string {
	path := f.Png
	if !utils.ExistsFiles(path) {
		path = f.Jpg
	}
	if !utils.ExistsFiles(path) {
		path = f.Path
	}
	return utils.ImageToString(path)
}

func (f Movie) GetPng() string {
	path := f.Path
	return utils.GetPng(path, "png")
}
