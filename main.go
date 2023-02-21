package main

import (
	"os"
	"zocket/articles"
	"zocket/task2"

	"github.com/gin-gonic/gin"
)

func main(){
    r := gin.New()

    r.POST("/article",articles.Add)
    r.GET("/article",articles.Get)
    r.GET("/article/:id",articles.GetOne)
    r.PUT("/article/:id",articles.Update)
    r.DELETE("/article/:id",articles.Delete)

    r.GET("/task2",task2.ReadFile)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    r.Run(":" + port)
}
