package main

import (
	"fmt"
	"github-actions-project/models"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var dbConfig struct {
	db *gorm.DB
}

func main() {

	// dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=require",
	// 	os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"),
	// 	os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))

	log.Printf("THE DATABASE URL IS ::: ", os.Getenv("POSTGRES_HOST"))
	// dns := os.Getenv("DATABASE_URL")

	// db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	// if err != nil {
	// 	log.Printf("ERROR IN CONNECTION ::::::: ", err.Error())
	// 	panic(err)
	// }
	router := gin.Default()
	router.GET("/v1/tasks", func(c *gin.Context) {
		var tasks []models.Task
		fmt.Print(tasks)
		// db.First(&tasks)
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "success",
			"data":    "tasks",
		})
	})
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "5000"
	}
	router.Run(":" + port)
}
