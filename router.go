package myrouter

import (
    "github.com/gin-gonic/gin"
)

func HomeHandler(context *gin.Context) {
    context.JSON(200, gin.H{"message": "Hello, World!"})
}

func UserHandler(context *gin.Context) {
    name := context.Param("name")
    context.JSON(200, gin.H{"message": "Hello, " + name + "!"})
}

func RegisterRoutes(router *gin.Engine) {
    router.GET("/", HomeHandler)
    router.GET("/user/:name", UserHandler)
}
