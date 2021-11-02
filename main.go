package main

import (
	"fmt"
	"github-actions-project/models"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConfig struct {
	db *gorm.DB
}

func main() {
	fmt.Println(os.Getenv("DB_USER"))
	fmt.Printf("POSTGRES USER IS %s", os.Getenv("DB_HOST"))

	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"), os.Getenv("SSL_MODE"))

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	router := gin.Default()
	router.GET("/v1/tasks", func(c *gin.Context) {
		var tasks []models.Task
		db.First(&tasks)
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "success",
			"data":    tasks,
		})
	})
	router.Run(":8080")

}
