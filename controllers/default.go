package controllers

import (
	model "crud/models"
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

// type MainController struct {
// 	beego.Controller
// }

type AddController struct {
	beego.Controller
}

func (this *AddController) Get() {
	this.TplName = "index.tpl"
}

func (this *AddController) Post() {
	o := orm.NewOrm()
	var user model.User
	user.Firstname = this.GetString("firstname")
	user.LastName = this.GetString("lastname")
	o.Insert(&user)

	id, err := o.Insert(&user)
	if err == nil {
		fmt.Println(id)
	}

}

// func (c *MainController) Get() {
// 	c.Data["Website"] = "beego.me"
// 	c.Data["Email"] = "astaxie@gmail.com"
// 	c.TplName = "index.tpl"
// }
