package main

import (
	"github-actions-project/models"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConfig struct {
	db *gorm.DB
}

func main() {

	// dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s",
	// 	os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"),
	// 	os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"), os.Getenv("SSL_MODE"))

	log.Printf("THE DATABASE URL IS ::: ", os.Getenv("DATABASE_URL"))
	dns := os.Getenv("DATABASE_URL")
	connection, _ := pq.ParseURL(dns)
	connection += " sslmode=require"

	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		log.Printf("ERROR IN CONNECTION ::::::: ", err.Error())
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
