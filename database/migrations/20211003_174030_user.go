package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type User_20211003_174030 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &User_20211003_174030{}
	m.Created = "20211003_174030"

	migration.Register("User_20211003_174030", m)
}

// Run the migrations
func (m *User_20211003_174030) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *User_20211003_174030) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
