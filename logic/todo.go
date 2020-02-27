package logic

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"gin_sql/dao"
)

// createTodo 方法添加一条新的 todo 数据
func CreateTodo(c *gin.Context) {
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	title := c.PostForm("title")
	err := dao.DoCreateData(title, completed)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "resourceId": err})
}

// fetchAllTodo 返回所有的 todo 数据
func FetchAllTodo(c *gin.Context) {
	todos, err := dao.DoFetchAllTodo()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": todos})
}

// fetchSingleTodo方法返回一条 todo 数据
func FetchSingleTodo(c *gin.Context) {
	todoID := c.Param("id")
	todo, err := dao.DoFetchSingleTodo(todoID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": todo})
}

// updateTodo 方法 更新 todo 数据
func UpdateTodo(c *gin.Context) {
	todoID := c.Param("id")
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	title := c.PostForm("title")
	err := dao.DoUpdateTodo(todoID, title, completed)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo updated successfully!"})
}

// deleteTodo 方法依据 id 删除一条todo 数据
func DeleteTodo(c *gin.Context) {
	todoID := c.Param("id")
	err := dao.DoDeleteTodo(todoID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully!"})
}
