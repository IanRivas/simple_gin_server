package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context){
	//enviamos un 200 y un json con el mensaje okey
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}