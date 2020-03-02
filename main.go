package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/SoftNotWare/study-ground/model"

	"github.com/SoftNotWare/study-ground/controller"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
)

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	fis, err := ioutil.ReadDir(templatesDir)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Fatal("加载html模板失败")
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
		log.WithFields(log.Fields{"name": name, "path": filepath.Join(templatesDir, fi.Name())}).Debug("增加模板")
	}

	return r
}

func main() {
	model.InitDB()

	log.SetLevel(log.DebugLevel)

	r := gin.New()
	// 安装logrus日志中间件
	r.Use(ginlogrus.Logger(log.New()), gin.Recovery())

	r.HTMLRender = loadTemplates("www")

	r.Static("js", "www/js")
	r.Static("css", "www/css")
	r.Static("image", "www/image")
	r.Static("fonts", "www/fonts")
	r.Static("danmu", "www/danmu")

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "home.html", nil)
	})

	r.GET("/play", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "play.html", nil)
	})

	r.POST("/login", controller.Register)

	log.Info("服务器启动成功")

	r.Run(":2233")
}
