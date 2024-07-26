package controllers

import "github.com/gin-gonic/gin"

type Controllers interface {
	InitRoutes(rg *gin.RouterGroup)
}
