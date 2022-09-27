package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
  	"gorm.io/gorm"
  	"fmt"
)

type Test struct {
  gorm.Model
  Name string
}

func main() {
	// database connection settings
	dsn := "root:Asdf1234!@tcp(127.0.0.1:3306)/retailer"
  	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

  	var test Test;
  	db.First(&test);

  	fmt.Println("test is ", test)

	// init setup for gin server
	r := gin.Default()
	r.GET("/hello", func(c * gin.Context) {
		c.JSON(200, gin.H {
			"message": "hello",
		})
	})
	r.Run()
}