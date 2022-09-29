package main

import (
	"gorm.io/driver/mysql"
  	"gorm.io/gorm"
  	"fmt"
)

type Product struct {
	Id string `gorm:"primary_key" json:"id"`
	Name string `json:"product_name"`
	Price float64 `json:"price" json:"type:decimal(10,2)"`
	Quantity int `json:"quantity"`
}

type Order struct {
	Id string `gorm:"primary_key" json:"id"`
	CustomerId string `json:"-" `
	ProductId string `json:"product_id"`
	Quantity int `json:"quantity"`
	Status string `json:"status "`
}

var db *gorm.DB;

func getDatabaseConnection() {
	dsn := "root:Asdf1234!@tcp(127.0.0.1:3306)/retailer"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if (err != nil) {
		fmt.Println("Database connection error: ", err)
	} 
	db = conn
}