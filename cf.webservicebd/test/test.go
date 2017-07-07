package main

import (
	"fmt"

	"../orm"
)

func main() {
	orm.CreateConnection()
	orm.DropTables()
	orm.CreateTables()
	var users orm.Users
	var user orm.User
	for i := 0; i < 1000; i++ {
		user.Create(fmt.Sprintf("%s%d", "user", i), fmt.Sprintf("%s%d", "password", i), fmt.Sprintf("%s%d%s", "user", i, "@gmail.com"))
	}
	fmt.Println(*users.FindAll())
	// orm.DropTables()
	orm.CloseConnection()
}

//gg
