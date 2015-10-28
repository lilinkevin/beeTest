package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplNames = "login.tpl"
}

func (this *LoginController) Post() {
	username := this.Input().Get("username")
	password := this.Input().Get("pwd")
	autoLogin := this.Input().Get("autoLogin") == "on"

	if beego.AppConfig.String("username") == username && beego.AppConfig.String("password") == password {
		fmt.Println("login suc")
		if autoLogin {
			maxAge := 1<<32 - 1
			this.Ctx.SetCookie("username", username, maxAge, "/")
			this.Ctx.SetCookie("password", password, maxAge, "/")
		}
		this.Redirect("/", 301)
	}
	return
}

func CheckLogin(ctx *context.Context) (bool, string) {
	val, err := ctx.Request.Cookie("username")
	if err != nil {
		return false, ""
	}
	username := val.Value

	val, err = ctx.Request.Cookie("password")
	if err != nil {
		return false, ""
	}
	password := val.Value

	return beego.AppConfig.String("username") == username && beego.AppConfig.String("password") == password, username

}
