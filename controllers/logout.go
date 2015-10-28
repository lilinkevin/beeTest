package controllers

import (
	"github.com/astaxie/beego"
)

type LogoutController struct {
	beego.Controller
}

func (this *LogoutController) Get() {
	this.Ctx.SetCookie("username", "", 0, "/")
	this.Ctx.SetCookie("password", "", 0, "/")
	this.Redirect("/", 301)
	return
}
