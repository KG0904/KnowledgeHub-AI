package services

import (
	"go-gateway/engine/handler"

	"github.com/gin-gonic/gin"
)

func AddServices(r *gin.Engine) {
	api := r.Group("/v1")
	{
		api.GET("/test", handler.TestHandler)

		api.POST("/document", handler.UploadDocument)

		//api.PATCH("/document/:id", handler.UpdateDocumentStatus)
	}
}
