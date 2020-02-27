package controller

import (
	"github.com/gin-gonic/gin"
)

func GinRouter(r *gin.Engine) *gin.Engine {
	//前缀为api/v1
	rr := r.Group("/api/v1")
	{
		Routers(rr)
	}
	return r
}
