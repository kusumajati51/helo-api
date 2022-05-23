package connection

import (
	"example/hello-api/config"
	"example/hello-api/models"
	"log"

	"github.com/gin-gonic/gin"
)

func GetConnection(ginEngine *gin.Engine) {
	db, err := config.GetDbConfig()
	if err != nil {
		log.Println("Connections")
	} else {
		log.Println("Connection Established")
	}
	db.AutoMigrate(models.User{})
	ginEngine.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
}
