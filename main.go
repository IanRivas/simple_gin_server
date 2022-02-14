package main

import (
	"github.com/IanRivas/simple_gin_server/controllers"
	"github.com/gin-gonic/gin"
) 


func main(){
	// basic gin middleware
	r := gin.Default()

	r.GET("/", controllers.Get)
	r.Run(":9090")

}