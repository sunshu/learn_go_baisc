package main

import (
	_ "go_web_page/routers" // 只调用init 方法
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"go_web_page/models"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:root@/orm_test?charset=utf8")
}

func main() {
	fmt.Println("Hello, World!")

	o := orm.NewOrm()
	o.Using("go_web_page") // 默认使用 default，你可以指定为其他数据库
	title :=  models.TITLE{}
	o.Read(title.TITLE_PARENT_ID)


	fmt.Println()
	beego.Run(":81")






}

