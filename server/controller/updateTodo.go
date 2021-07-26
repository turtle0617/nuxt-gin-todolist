package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/turtle0617/nuxt-gin-todolist/server/config"
	"github.com/turtle0617/nuxt-gin-todolist/server/model"
)

func UpdateTodo(c *gin.Context){
	var todo model.Todo
	var payload model.Todo
	id := c.Param("id")
	if err :=  config.GetDB().Where("id = ?",id).First(&todo).Error; err != nil {
		c.JSON(401, gin.H{"error":err})
	}

	if err:= c.ShouldBindJSON(&payload); err!=nil {
		c.JSON(401, gin.H{"error": err.Error()})
	}

	if err:= config.GetDB().Model(&todo).Updates(payload).Error; err!= nil {
		c.JSON(401, gin.H{"error": err})
	}

	c.JSON(200, gin.H{
		"data": todo,
	})
}