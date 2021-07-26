package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/turtle0617/nuxt-gin-todolist/server/config"
	"github.com/turtle0617/nuxt-gin-todolist/server/controller"
)

func main() {
	config.GetEnv()
	r := gin.Default()

	corsConfig := config.CorsConfig()
	r.Use(cors.New(corsConfig))

	r.GET("/todos",controller.GetTodos)

	r.POST("/todos",controller.AddTodo)

	r.PATCH("/todos/:id",controller.UpdateTodo)

	r.DELETE("/todos/:id",controller.DeleteTodo)
	r.Run()
}
