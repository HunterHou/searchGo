package datamodels

import (
	"fmt"
	"time"

	"../utils/fileUtils"
)

//声明一个File结构体 表示一个文件信息
type File struct {
	Code     string "json:'code'"
	Name     string
	Path     string
	Actress  string
	FileType string
	DirPath  string
	Size     int64
	SizeStr  string
	CTime    string
	MTime    string
}

func NewFile(dir string, path string, name string, fileType string, size int64, modTime time.Time) File {
	// 使用工厂模式 返回一个 File 实例
	result := File{
		Code:     fileUtils.GetCode(name),
		Name:     name,
		Path:     path,
		Actress:  fileUtils.GetCode(name),
		FileType: fileType,
		DirPath:  dir,
		Size:     size,
		SizeStr:  fileUtils.GetSizeStr(size),
		CTime:    "",
		MTime:    modTime.Format("2006-01-02 15:04:05"),
	}
	return result
}

func (f File) GetFileInfo() string {
	//
	info := fmt.Sprintf("name: %v\t code:%v\t fileType:%v\t sizeStr:%v\t actress:%v\t path:%v\t",
		f.Name, f.Code, f.FileType, f.SizeStr, f.Actress, f.Path)
	return info
}