package main

import (
	"github.com/crision98/lisander-golang-backend/controller"
	"github.com/crision98/lisander-golang-backend/database"
)

func main() {
	if database.CheckConnection() == 0 {
		return
	}
	controller.Controllers()
}
