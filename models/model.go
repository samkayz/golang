package model

import (
	"github.com/beego/beego/v2/client/orm"
)

type User struct {
	Id        int
	Firstname string
	LastName  string

	// Profile *Profile `orm:"rel(one)"` // OneToOne relation
}

type Logs struct {
	Id      int
	Loguser string
	Name    string
}

type Profile struct {
	Id   int
	Age  int16
	User *User `orm:"reverse(one)"` // Reverse relationship (optional)
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(User), new(Profile), new(Logs))
}
