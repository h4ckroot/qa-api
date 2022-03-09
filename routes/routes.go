package routes

import (
	"github.com/gin-gonic/gin"
	"shamilabs.com/qa-api/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.POST("question", controllers.CreateQuestion)
		v1.PUT("question/:id/answer", controllers.CreateAnswer)
		v1.PUT("question/:id/tag", controllers.UpdateQuestionWithTag)
		// v1.PUT("question/:id/comment", controllers.UpdateQuestionWithComment)
		v1.DELETE("question/:id", controllers.DeleteQuestion)
		v1.GET("question/:id", controllers.GetQuestion)
		v1.GET("healthz", controllers.IsHealthy)
		v1.GET("is-ready", controllers.IsReady)
	}

	return r
}
