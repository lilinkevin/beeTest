package main

import (
	"beeTest/models"
	_ "beeTest/routers"
	_ "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	o := orm.NewOrm()
	orm.Debug = true
	orm.RunSyncdb("default", false, true)

	o.Using("default")
	user := models.User{Name: "李琳"}
	o.Insert(&user)
	beego.Run()

}

const (
	_DIRER_NAME = "mysql"
)

func init() {
	orm.RegisterDriver(_DIRER_NAME, orm.DR_MySQL)
	orm.RegisterDataBase("default", _DIRER_NAME, "root:780810@tcp(192.168.10.22:33306)/test?charset=utf8")
}
