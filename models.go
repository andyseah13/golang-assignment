package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
*********************

	Object Models

*********************
*/
type Product struct {
	Id       string  `gorm:"primary_key;<-:false" json:"id"`
	Name     string  `json:"product_name"`
	Price    float64 `json:"price" json:"type:decimal(10,2)"`
	Quantity int     `json:"quantity"`
}

type Order struct {
	Id         string `gorm:"primary_key;<-:false" json:"id"`
	CustomerId string `json:"customer_id" `
	ProductId  string `json:"product_id"`
	Quantity   int    `json:"quantity"`
	Status     string `json:"status "`
}

/*
********************

	Response Objects

*********************
*/
type ProductListResponse struct {
	Products []Product `json:"products"`
}

var db *gorm.DB

func getDatabaseConnection() {
	dsn := "root:Asdf1234!@tcp(127.0.0.1:3306)/retailer"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Database connection error: ", err)
	}
	db = conn
}
