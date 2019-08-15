package routers

import (
	"go_web_page/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	//beego.Router("/index_body", &controllers.IndexController{})
}
