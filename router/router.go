package router

import (
	"eyeGlassesNew/controller"
	"github.com/beego/beego/v2/server/web"
)

func init() {
	// 使用命名空间注册路由
	//ns := web.NewNamespace("v1",
	//	web.NSBefore(func(ctx *context.Context) {
	//		fmt.Println("过滤器之前")
	//	}),
	//	web.NSAfter(func(ctx *context.Context) {
	//		fmt.Println("过滤器之后")
	//	}),
	//	web.NSGet("/login", controller.Login),
	//	web.NSGet("/register", controller.Register),
	//)
	//web.AddNamespace(ns)

	// 原始的登录注册相关的路由
	//web.Get("/login", controller.Login)
	//web.Get("/register", controller.Register)
	//fmt.Println("路由注册完成")
	web.CtrlGet("/login", (*controller.UserController).Login)
	web.CtrlGet("/register", (*controller.UserController).Register)
	web.CtrlPost("/login", (*controller.UserController).LoginPost)
	web.CtrlGet("/index", (*controller.IndexController).ShowIndex)
	web.CtrlPost("/register", (*controller.UserController).RegisterPost)
}
