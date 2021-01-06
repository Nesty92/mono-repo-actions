package main

import "github.com/gin-gonic/gin"

func main() {
    router := gin.Default()

    router.GET("/posts", func(c *gin.Context) {
        c.String(200, "posts")
    })

    router.Run()
}