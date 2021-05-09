package main

import (
	"fmt"

	"todo_app1/app/controllers"
	"todo_app1/app/models"
)

func TestConnection() {

}

func main() {
	fmt.Println(models.Db)
	go controllers.StartMainServer()

	for {

	}
}
