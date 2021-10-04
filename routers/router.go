package routers

import (
	"crud/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// beego.Router("/", &controllers.MainController{})
	beego.Router("/", &controllers.AddController{})
}
