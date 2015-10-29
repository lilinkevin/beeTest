package controllers

import (
	"beeTest/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type CatagoryController struct {
	beego.Controller
}

func (this *CatagoryController) Get() {
	if hasLogin, username := CheckLogin(this.Ctx); hasLogin {
		this.Data["HasLogin"] = hasLogin
		this.Data["UserName"] = username
		this.TplNames = "catagory.tpl"
		this.Data["IsCatagory"] = true

		o := orm.NewOrm()
		var catagroys []*models.Category
		o.QueryTable("Category").All(&catagroys)
		fmt.Println(catagroys)
		this.Data["Catagorys"] = catagroys

	} else {
		this.Redirect("/login", 301)
	}

}

func (this *CatagoryController) Post() {

	if isLogin, _ := CheckLogin(this.Ctx); isLogin {
		title := this.Input().Get("title")
		catagroy := models.Category{Title: title, Ctime: time.Now(), Views: 0, TopicTime: time.Now(), TopicUserId: 1}

		o := orm.NewOrm()
		o.Insert(&catagroy)
		this.Redirect("/catagory", 301)

	} else {
		this.Redirect("/login", 301)
	}

}
