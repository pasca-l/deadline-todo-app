package main

import (
	"fmt"
	// "log"
	// "app/config"
	"app/models"
)

func main() {
	// fmt.Println(config.Config)
	// fmt.Println(config.Config.Srv.Port)

	// log.Println("test")

	fmt.Println("at main.go", models.Db)

	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@example.com"
	// u.Password = "password"
	// fmt.Println(u)

	// u.CreateUser()

	u, _ := models.GetUser(2)
	fmt.Println(u)
}
