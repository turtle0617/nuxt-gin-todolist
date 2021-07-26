package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/turtle0617/nuxt-todolist/server/config"
	"github.com/turtle0617/nuxt-todolist/server/model"
)

func GetTodos(c *gin.Context) {
	var todos []model.Todo
	config.GetDB().Find(&todos)
	c.JSON(200, gin.H{
		"data": todos,
	})
}