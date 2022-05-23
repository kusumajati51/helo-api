package router

import (
	"example/hello-api/controllers"
	"example/hello-api/router/middlewares"

	"github.com/gin-gonic/gin"
)

func GetRouter(routerEngine *gin.Engine) {
	routerEngine.GET("/user", controllers.FindUsers)
	routerEngine.POST("/user", controllers.CreateUser)
	routerEngine.POST("/login", controllers.UserLogin)
	secured := routerEngine.Group("").Use(middlewares.JwtAuthMiddleware())
		{
			secured.GET("/ping", controllers.GetDataUserLogin)
		}
	
}
