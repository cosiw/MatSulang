package router

import (
	"Back/pkg/api"

	"github.com/gin-gonic/gin"
)

func Router(apis *api.APIs) *gin.Engine {
	r := gin.Default()

	r.POST("/api/login", apis.Login)
	r.POST("/api/user", apis.InsertUser)
	r.DELETE("/api/user/:id", apis.DeleteUser)
	r.PATCH("/api/user/:id", apis.UpdateUser)
	return r
}
