package main

import (
	_ "eyeGlassesNew/router"
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/server/web"
)

func main() {
	staticHome, _ := config.String("staticHome")
	web.SetStaticPath("/static", staticHome)
	web.Run()
}
