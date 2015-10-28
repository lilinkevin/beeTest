package controllers

import (
	_ "beeTest/models"
	_ "fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	hasLogin, username := CheckLogin(c.Ctx)
	c.Data["HasLogin"] = hasLogin
	c.Data["UserName"] = username
	c.TplNames = "index.tpl"
}
