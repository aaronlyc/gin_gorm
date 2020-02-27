package controller

import (
	"github.com/gin-gonic/gin"

	"gin_sql/logic"
)

func Routers(r *gin.RouterGroup) {
	v1 := r.Group("/todos")
	{
		v1.POST("/", logic.CreateTodo)
		v1.GET("/", logic.FetchAllTodo)
		v1.GET("/:id", logic.FetchSingleTodo)
		v1.PUT("/:id", logic.UpdateTodo)
		v1.DELETE("/:id", logic.DeleteTodo)
	}
	return
}
