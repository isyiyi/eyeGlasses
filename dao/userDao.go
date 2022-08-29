package dao

import (
	"crypto/sha256"
	"eyeGlassesNew/model"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

func init() {
	mysql, _ := config.GetSection("mysql")
	orm.RegisterModel(new(model.Users))
	err := orm.RegisterDataBase("default", "mysql", mysql["username"]+":"+mysql["password"]+"@tcp("+mysql["ip"]+")/eyeGlasses?charset=utf8")
	if err != nil {
		fmt.Println("注册数据库失败:", err)
	}
}

func Login(user *model.Users) bool {
	o := orm.NewOrm()
	user.Password = fmt.Sprintf("%x", sha256.Sum256([]byte(user.Password)))
	o.Raw("select `id` from users where username = ? and `password` = ?", user.UserName, user.Password).QueryRow(&(user.Id))
	if user.Id != 0 {
		err := o.Read(user)
		if err != nil {
			fmt.Println(err)
			return false
		} else {
			return true
		}
	} else {
		return false
	}
}

func Register(user *model.Users) bool {
	o := orm.NewOrm()
	user.Password = fmt.Sprintf("%x", sha256.Sum256([]byte(user.Password)))
	res, err := o.Raw("insert into users (username, `password`, is_delete, create_time) values (?, ?, 0, now())", user.UserName, user.Password).Exec()
	if err != nil {
		if strings.Contains(fmt.Sprintf("%s", err), "Duplicate entry") {
			fmt.Println("用户名已被占用")
			return false
		}
		fmt.Println(err)
		return false
	}
	id, _ := res.LastInsertId()
	fmt.Println(id)
	return true
}
