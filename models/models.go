package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id   int64
	Name string
}

type Category struct {
	Id          int64
	Title       string
	Ctime       time.Time `orm:"index"`
	Views       int64
	TopicTime   time.Time `orm:"index"`
	TopicUserId int64
}

type Topic struct {
	Id               int64
	UserId           int64
	Title            string
	Content          string `orm:"size(5000)"`
	Attachment       string
	CTime            time.Time
	UTime            time.Time
	Views            int64
	Author           string
	ReplayTime       time.Time
	ReplayLastUserId int64
	ReplayCount      int64
}

type School struct {
	Id   int64
	Name string `orm:"size(50)"`
}

func init() {
	orm.RegisterModel(new(User), new(School), new(Category), new(Topic))
}
