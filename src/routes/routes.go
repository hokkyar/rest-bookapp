package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hokkyar/rest-bookapp/src/controllers"
)

func InitRoutes(r *gin.Engine) {

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.GET("/users", controllers.GetUsers)
	r.GET("/users/:id", controllers.GetUserById)
	r.POST("/users", controllers.PostUser)
	r.PUT("/users/:id", controllers.PutUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	r.GET("/books", controllers.GetBooks)
	r.GET("/books/:id", controllers.GetBookById)
	r.POST("/books", controllers.PostBook)
	r.PUT("/books/:id", controllers.PutBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

}
