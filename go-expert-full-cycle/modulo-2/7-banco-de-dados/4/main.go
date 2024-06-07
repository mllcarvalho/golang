package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID           int `gorm:"primary_key"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type Category struct {
	ID   int `gorm:"primary_key"`
	Name string
}

type SerialNumber struct {
	ID        int `gorm:"primary_key"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	//create category
	db.Create(&Category{
		Name: "Eletrônicos",
	})

	//create product
	db.Create(&Product{
		Name:       "Notebook",
		Price:      1899.90,
		CategoryID: 1,
	})

	//create serial number
	db.Create(&SerialNumber{
		Number:    "123456",
		ProductID: 1,
	})

	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
	}

}
