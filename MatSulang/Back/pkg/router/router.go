package router

import (
	"Back/pkg/api"

	"github.com/gin-gonic/gin"
)

func Router(apis *api.APIs) *gin.Engine {
	r := gin.Default()

	r.POST("/login", apis.Login)
	return r
}
