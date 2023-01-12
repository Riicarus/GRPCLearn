package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitAPI() {
	router := gin.Default()

	apiV1 := router.Group("/api/v1") 
	{
		apiV1.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "OK",
			})
		})
	}

	router.Run()
}