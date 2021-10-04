package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Logs_20211003_175253 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Logs_20211003_175253{}
	m.Created = "20211003_175253"

	migration.Register("Logs_20211003_175253", m)
}

// Run the migrations
func (m *Logs_20211003_175253) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *Logs_20211003_175253) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
