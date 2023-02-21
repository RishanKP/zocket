package articles

import (
	"strconv"
    "errors"

	"github.com/gin-gonic/gin"
)

type Article struct{
    Id int `json:"id"`
    Author string `json:"author"`
    Title string `json:"title"`
    Description string `json:"description"`         
}

var Articles []Article
var i = 0

func GetPosition(id int) (int,error){
    for j:=0;j<len(Articles);j++ {
        if Articles[j].Id == id{
            return j,nil
        }
    }

    return -1,errors.New("not found")
}

func DeleteArticle(pos int){
    Articles = append(Articles[:pos],Articles[pos+1:]... )
    return
}

func Add(c *gin.Context){
    var a Article
    c.BindJSON(&a)

    i = i+1
    a.Id = i
    Articles = append(Articles,a)

    c.JSON(200,gin.H{
        "message":"article inserted",
    })
}

func Get(c *gin.Context){
    if len(Articles) == 0{
        c.JSON(200,gin.H{
            "message":"no articles found",
        })

        return
    }
    c.JSON(200,gin.H{
        "message":"fetched articles",
        "articles": Articles,
    })
}

func GetOne(c *gin.Context){
    id := c.Param("id")

    aid,err := strconv.Atoi(id)
    if err !=nil {
        c.JSON(400,gin.H{
            "message":"invalid id",
        })
        
        return
    }
    
    pos,err := GetPosition(aid)
    if err != nil{
         c.JSON(400,gin.H{
            "message":"article not found",
        })

        return
    }

    c.JSON(200,gin.H{
        "message":"fetched articles",
        "article": Articles[pos],
    })
}

func Delete(c *gin.Context){
    id := c.Param("id")
    
    aid,err := strconv.Atoi(id)
    if err !=nil {
        c.JSON(400,gin.H{
            "message":"invalid id",
        })
        
        return
    }
    
    pos,err := GetPosition(aid)
    if err != nil{
         c.JSON(400,gin.H{
            "message":"article not found",
        })

        return
    }

    DeleteArticle(pos)
    c.JSON(200,gin.H{
        "message":"article deleted",
    })
}

func Update(c *gin.Context){
    id := c.Param("id")
    var a Article

    c.BindJSON(&a)

    aid,err := strconv.Atoi(id)
    if err !=nil {
        c.JSON(400,gin.H{
            "message":"invalid id",
        })
        
        return
    }
    
    pos,err := GetPosition(aid)
    if err != nil{
         c.JSON(400,gin.H{
            "message":"article not found",
        })

        return
    }

    Articles[pos].Title = a.Title
    Articles[pos].Author = a.Author
    Articles[pos].Description = a.Description

    c.JSON(200,gin.H{
        "message":"article updated",
    })
}
