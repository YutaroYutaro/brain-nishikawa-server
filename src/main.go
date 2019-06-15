package main // import "src"
import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"src/app/interface/api"
	"src/app/registry"
)

type Meeting struct {
	ID    int
	Title string
}

func main() {
	ctn, err := registry.NewContainer()
	if err != nil {
		log.Fatal(err.Error())
	}
	r := gin.Default()

	db, err := gorm.Open("mysql", "root:password@tcp(db-server:3306)/test_db")
	defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	api.Api(ctn, r, db)
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(
	//		200,
	//		gin.H{"message": "pong"},
	//	)
	//})
	//
	//db, err := gorm.Open("mysql", "root:password@tcp(db-server:3306)/test_db")
	//defer db.Close()
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//
	//r.GET("/meeting", func(c *gin.Context) {
	//	var meetings []Meeting
	//	db.Table("meeting").Select("id, title").Find(&meetings)
	//
	//	c.JSON(
	//		200,
	//		meetings,
	//	)
	//})
	//
	//r.GET("meeting/:id", func(c *gin.Context) {
	//	var meeting Meeting
	//	id := c.Param("id")
	//
	//	db.Table("meeting").Select("id, title").Where("id = ?", id).Find(&meeting)
	//	c.JSON(
	//		200,
	//		meeting,
	//	)
	//})

	r.Run(":8080")
}
