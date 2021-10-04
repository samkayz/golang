package main

import (
	// model "crud/models"
	_ "crud/routers"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:victoria1992@/crud")
}

func main() {
	// o := orm.NewOrm()

	// profile := new(model.Profile)
	// profile.Age = 30

	// user := new(model.User)
	// user.Profile = profile
	// user.Name = "slene"

	// fmt.Println(o.Insert(profile))
	// fmt.Println(o.Insert(user))

	// Database alias.
	// name := "default"

	// // Drop table and re-create.
	// force := true

	// // Print log.
	// verbose := true

	// // Error.
	// err := orm.RunSyncdb(name, force, verbose)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	beego.Run()
}
