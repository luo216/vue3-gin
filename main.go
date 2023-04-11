package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    "github.com/luo216/vue3-gin/myrouter"
)

func main() {
    router := gin.Default()

    // 添加跨域中间件
    router.Use(cors.Default())

    myrouter.RegisterRoutes(router)
    router.Run(":1234")
}
