package task2

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
    "strconv"

	"github.com/gin-gonic/gin"
)

type Team struct{
    Name string `json:"name"`
    Points int `json:"points"`
}

type PointsTable struct{
    Teams []Team `json:"pointsTable"`
}
func ReadFile(c *gin.Context){
    f,err := os.Open("/etc/secrets/file.csv")
    if err != nil{
        c.JSON(400,gin.H{
            "message":"error opening file",
            "error":err,
        })
        return
    }

    defer f.Close()

    csvReader := csv.NewReader(f)

    var t PointsTable
    
    //skip first line
    rec,err := csvReader.Read()

    for{    
        rec,err = csvReader.Read()
        if err == io.EOF{
            break
        }

        fmt.Println(rec)
        var temp Team
        for j,field := range rec{
            if j==0{
               temp.Name = field 
            }else{
               temp.Points, _  = strconv.Atoi(field) 
            }
        }

        t.Teams = append(t.Teams, temp)
    }

    c.JSON(200,gin.H{
        "message":"file read successfully",
        "data": t,
    })
}
