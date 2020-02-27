package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"gin_sql/controller"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	gin.ForceConsoleColor()
	r := gin.Default()
	controller.GinRouter(r)
	r.Run(":8090")
}
