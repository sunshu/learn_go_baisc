package main

import (
	"fmt"
	"github.com/astaxie/beego"
)

// 普通
type RESTfulController struct {
	beego.Controller
}

func (this *RESTfulController) Get() {
	this.Ctx.WriteString("Hello World in GET Method")
}

func (this *RESTfulController) Post() {
	this.Ctx.WriteString("Hello World in POST Method")
}

// 正则
type RegExpController struct {
	beego.Controller
}

func (this *RegExpController) Get() {
	this.Ctx.WriteString(fmt.Sprintln("In RegExp Mode"))

	id := this.Ctx.Input.Param(":id")
	fmt.Println("id is: %s\n", id)
}

func main() {
	beego.Router("/RESTful", &RESTfulController{})

	// 正则路由 从PATH 中提取参数
	beego.Router("/RegExp1/?:id", &RegExpController{})
	beego.Router("RegExp2/:id([0-9]+)", &RegExpController{})
	beego.Router("RegExp3/:id([\\w]+)", &RegExpController{})

	// 服务启动 ip port
	beego.Run("127.0.0.1:8081")

}
