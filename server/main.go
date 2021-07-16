package main

import (
	"fmt"

	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
type Todo struct {
	ID    			uint `gorm:"primaryKey" json:"id"`
	Title 			string `json:"title"`
	Status      string `json:"status"`
}

func getEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	r := gin.Default()
	db := openDB()

	r.GET("/todos", func(c *gin.Context) {
		var todos []Todo
		db.Find(&todos)
		c.JSON(200, gin.H{
			"data": todos,
		})
	})

	r.POST("/todos",func(c *gin.Context){
		var payload Todo
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			return
		}
		
		createTodoResult := db.Select("Title").Create(&payload)
		searchTodoResult := db.First(&payload)

		if createTodoResult.Error != nil {
			c.JSON(401, gin.H{"error":createTodoResult.Error})
		}
		if searchTodoResult.Error != nil {
			c.JSON(401, gin.H{"error":searchTodoResult.Error})
		}

		c.JSON(200, gin.H{
			"data": payload,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func openDB()*gorm.DB{
	getEnv()
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"), os.Getenv(("POSTGRES_HOST_PORT")))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }
	return db
}