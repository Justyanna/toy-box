package runtime

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	service := NewRuntimeService()
	controller := NewRuntimeController(service)

	r.GET("/", controller.GetGameContext)
	r.POST("/", controller.InvokeMethod)
}
