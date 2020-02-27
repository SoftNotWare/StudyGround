package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/siddontang/go/log"
	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"

	"os"
)

func setup(r *gin.Engine) {
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	logrus.SetLevel(logrus.WarnLevel)

	// 安装logrus日志中间件
	log := logrus.New()
	log.SetLevel(logrus.WarnLevel)
	r.Use(ginlogrus.Logger(log), gin.Recovery())
}

func main() {
	r := gin.New()

	setup(r)

	r.Static("js", "www/js")
	r.Static("css", "www/css")
	r.Static("danmu", "www/danmu")

	var err error
	tpl := template.New("index")
	tpl, err = tpl.ParseFiles("www/index.html")
	if err != nil {
		log.Fatal("载入index模板失败：", err)
	}
	r.SetHTMLTemplate(tpl)
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	log.Info("服务器启动成功")

	r.Run(":2233")
}
