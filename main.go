package main

import (
	"os"
	"github.com/gin-gonic/gin"
    "zocket/articles"
)

func main(){
    r := gin.New()

    r.POST("/article",articles.Add)
    r.GET("/article",articles.Get)
    r.GET("/article/:id",articles.GetOne)
    r.DELETE("/article/:id",articles.Delete)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    r.Run(":" + port)
}
