package main

import (
	controllers "goproject/controllers"
	"goproject/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := models.SetupModels() // new
	// Provide db variable to controllers
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	r.GET("/users", controllers.FindUsers)
	r.POST("/users", controllers.CreateUser)       // create
	r.GET("/users/:id", controllers.FindUser)      // find by id
	r.PATCH("/users/:id", controllers.UpdateUser)  // update by id
	r.DELETE("/users/:id", controllers.DeleteUser) // delete by id
	r.POST("/login", controllers.LoginService)
	//r.POST("/login",controllers.LoginService)//login
	r.Run()
}
