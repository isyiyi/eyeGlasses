package controller

import (
	"eyeGlassesNew/dao"
	"eyeGlassesNew/model"
	"fmt"
	"github.com/beego/beego/v2/server/web"
)

type UserController struct {
	web.Controller
}

func (u *UserController) Login() {
	u.TplName = "login.html"
}

func (u *UserController) LoginPost() {
	user := model.Users{}
	user.UserName = u.GetString("username")
	user.Password = u.GetString("password")
	user.IsDelete = 0
	if dao.Login(&user) {
		u.Redirect("/index", 200)
		fmt.Println("登录成功")
	}
}

func (u *UserController) Register() {
	u.TplName = "register.html"
}

func (u *UserController) RegisterPost() {
	user := model.Users{}
	user.UserName = u.GetString("username")
	user.Password = u.GetString("password")
	if dao.Register(&user) {
		u.Redirect("/login", 200)
	} else {
		u.Redirect("/register", 200)
	}
}
