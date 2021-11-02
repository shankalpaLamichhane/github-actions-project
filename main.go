package main

import (
	"github-actions-project/models"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbConfig struct {
	db *gorm.DB
}

func main() {
	// dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASS"),
	// 	os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	dns := "root:password@tcp(127.0.0.1:3306)/actions?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to db")
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
