package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/turtle0617/nuxt-todolist/server/config"
	"github.com/turtle0617/nuxt-todolist/server/model"
)

func AddTodo(c *gin.Context){
	var payload model.Todo
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}
	
	createTodoResult := config.GetDB().Select("Title").Create(&payload)
	searchTodoResult := config.GetDB().First(&payload)

	if createTodoResult.Error != nil {
		c.JSON(401, gin.H{"error":createTodoResult.Error})
	}
	if searchTodoResult.Error != nil {
		c.JSON(401, gin.H{"error":searchTodoResult.Error})
	}

	c.JSON(200, gin.H{
		"data": payload,
	})
}