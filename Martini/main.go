package main

import (
	"Martini/controllers"

	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	m := martini.Classic()

	m.Group("/user", func(r martini.Router) {
		r.Get("/", controllers.GetAllUsers)
		r.Post("/", controllers.InsertUser)
		r.Put("/:idUser", controllers.UpdateUser)
		r.Delete("/:idUser", controllers.DeleteUser)
	})

	m.RunOnAddr(":8080")
}
