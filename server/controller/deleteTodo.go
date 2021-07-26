package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/turtle0617/nuxt-gin-todolist/server/config"
	"github.com/turtle0617/nuxt-gin-todolist/server/model"
)

func DeleteTodo(c *gin.Context){
	var todo model.Todo
	id := c.Param("id")

	if err := config.GetDB().Delete(&todo,"id = ?",id).Error; err != nil {
		c.JSON(401, gin.H{"error":err})
	}else {
		c.JSON(200, gin.H{
			"message": "Deleted successfully",
		})
	}
}