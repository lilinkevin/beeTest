package controllers

import (
	"github.com/astaxie/beego"
)

type CatagoryController struct {
	beego.Controller
}

func (this *CatagoryController) Get() {
	this.TplNames = "catagory.tpl"
}
