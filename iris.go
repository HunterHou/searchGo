package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/mvc"
)
import (
	"./cons"
	"./controller"
	"./service"
	"./utils"
	"path/filepath"
)

var curDir string
var staticDir string

// 打包命令
// 1 命令行UI 常规打包 go build
// 2 无窗口  go build -o app/app.exe -ldflags "-H=windowsgui"

func init() {
	curDir, _ := filepath.Abs(".")
	//if !strings.HasSuffix(curDir, "src") {
	//	curDir += "/src"
	//}
	cons.DirFile = curDir + "\\dirList.ini"
	dict := service.ReadDictionary(cons.DirFile)
	dirs := dict.GetProperty("dir")
	baseUrls := dict.GetProperty("BaseUrl")
	if len(baseUrls) > 0 {
		cons.BaseUrl = baseUrls[0]
	}
	Images := dict.GetProperty("Images")
	if len(Images) > 0 {
		cons.Images = Images
	}
	VideoTypes := dict.GetProperty("VideoTypes")
	if len(VideoTypes) > 0 {
		cons.VideoTypes = VideoTypes
	}
	Docs := dict.GetProperty("Docs")
	if len(Images) > 0 {
		cons.Docs = Docs
	}

	cons.SetBaseDir(dirs)

	staticDir = curDir + "/static"
	cons.Play = utils.ImageToString(staticDir + "/image/play.jpg")
	cons.Open = utils.ImageToString(staticDir + "/image/open.jpg")
	cons.Change = utils.ImageToString(staticDir + "/image/change.jpg")
	cons.Replay = utils.ImageToString(staticDir + "/image/replay.jpg")
	cons.Close = utils.ImageToString(staticDir + "/image/close.jpg")
	cons.Stop = utils.ImageToString(staticDir + "/image/stop.jpg")
}

func main() {
	app := iris.New()
	app.Handle("GET", "/", func(ctx iris.Context) {
		app.Logger().Info(ctx.Path())
		ctx.HTML("<h1>hello world!!!</h1>")

	})
	//done := make(chan bool, 1)
	//quit := make(chan os.Signal, 1)
	//signal.Notify(quit,os.Interrupt)
	customLogger := logger.New(logger.Config{
		Status:             true,
		IP:                 true,
		Method:             true,
		Path:               true,
		Query:              false,
		Columns:            false,
		MessageContextKeys: nil,
		MessageHeaderKeys:  nil,
		LogFunc:            nil,
		LogFuncCtx:         nil,
		Skippers:           nil,
	})
	app.Use(customLogger)
	app.RegisterView(iris.Django(staticDir, ".html"))
	app.HandleDir("/", staticDir)
	// http.Handle("/",http.FileServer(AssetFS()))
	app.Logger().SetLevel("debug")
	mvc.New(app).Handle(new(controller.TestController))
	mvc.New(app).Handle(new(controller.FileController))
	mvc.New(app).Handle(new(controller.SettingController))
	//启动直播流
	//liveGo := staticDir+ "/livego.exe"
	//utils.ExecCmdStart(liveGo)
	utils.ExecCmdStart("http://127.0.0.1:80/views")
	app.Run(iris.Addr(":80"), iris.WithConfiguration(iris.Configuration{
		DisableStartupLog:    false,
		FireMethodNotAllowed: false,
		TimeFormat:           "2019-11-10 18:10:33",
		Charset:              "uft-8",
	}))

}
