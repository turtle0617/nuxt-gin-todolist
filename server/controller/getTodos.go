package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/turtle0617/nuxt-gin-todolist/server/config"
	"github.com/turtle0617/nuxt-gin-todolist/server/model"
)

func GetTodos(c *gin.Context) {
	var todos []model.Todo
	config.GetDB().Order("id asc").Find(&todos)
	c.JSON(200, gin.H{
		"data": todos,
	})
}