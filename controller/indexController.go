package controller

import "github.com/beego/beego/v2/server/web"

type IndexController struct {
	web.Controller
}

func (i *IndexController) ShowIndex() {
	i.TplName = "index.html"
}
