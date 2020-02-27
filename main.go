package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"

	"os"
)

func setup(r *gin.Engine) {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	log.SetLevel(log.WarnLevel)

	// 安装logrus日志中间件
	l := log.New()
	l.SetLevel(log.WarnLevel)
	r.Use(ginlogrus.Logger(l), gin.Recovery())
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	fis, err := ioutil.ReadDir(templatesDir)
	if err != nil {
		log.Fatal("加载html模板失败：", err)
	}

	for _, fi := range fis {
		if fi.IsDir() || filepath.Ext(fi.Name()) != ".html" {
			continue
		}

		name := filepath.Base(fi.Name())
		tpl := template.New(name)
		tpl.Delims("{#", "#}")
		tpl = template.Must(tpl.ParseFiles(filepath.Join(templatesDir, fi.Name())))
		r.Add(name, tpl)
	}

	return r
}

func main() {

	r := gin.New()
	setup(r)
	r.HTMLRender = loadTemplates("www")

	r.Static("js", "www/js")
	r.Static("css", "www/css")
	r.Static("image", "www/image")
	r.Static("fonts", "www/fonts")
	r.Static("danmu", "www/danmu")

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/play", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "play.html", nil)
	})

	log.Info("服务器启动成功")

	r.Run(":2233")
}
