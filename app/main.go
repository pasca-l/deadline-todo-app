package main

import (
	"fmt"

	"app/models"
	"app/controllers"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartServer()
}
