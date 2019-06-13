package main // import "src"
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

type Meeting struct {
	ID    int
	Title string
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(
			200,
			gin.H{"message": "pong"},
		)
	})

	db, err := gorm.Open("mysql", "root:password@tcp(db-server:3306)/test_db")
	defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	r.GET("/meeting", func(c *gin.Context) {
		var meetings []Meeting
		db.Table("meeting").Select("id, title").Find(&meetings)

		c.JSON(
			200,
			meetings,
		)
	})

	r.GET("meeting/:id", func(c *gin.Context) {
		var meeting Meeting
		id := c.Param("id")

		db.Table("meeting").Select("id, title").Where("id = ?", id).Find(&meeting)
		c.JSON(
			200,
			meeting,
		)
	})

	r.Run(":8080")
}
