package models

import (
	"github.com/astaxie/beego/orm"

)

type TITLE struct {
	ID					 string `orm:"column(uid);pk"`
	TITLE     			 string
	TITLE_ICON      	 string
	TITLE_ID        	 string
	TITLE_PARENT_ID 	 string
	TITLE_DESC      	 string
	LEVEL				 string
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(TITLE))
}

