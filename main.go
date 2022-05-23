package main

import (
	"example/hello-api/connection"
	"example/hello-api/router"
	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)


func main() {
	r := gin.Default()
	// Listen and Server in 0.0.0.0:8080
	connection.GetConnection(r)
	router.GetRouter(r)
	r.Run(":8080")

}
