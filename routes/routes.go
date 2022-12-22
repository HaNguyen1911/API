package routes

import (
	"API/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {

	r := gin.Default()
	v1 := r.Group("/")
	{
		v1.POST("api", controllers.Create_Form)
		v1.GET("api", controllers.GetALL_Form)
		v1.GET("api/:id", controllers.GetById_Form)
		v1.PUT("api/:id", controllers.Update_Form)
		v1.DELETE("api/:id", controllers.Delete_Form)
	}
	return r
}
