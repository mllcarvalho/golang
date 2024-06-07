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
	ID       int `gorm:"primary_key"`
	Name     string
	Products []Product
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
	// db.Create(&Category{
	// 	Name: "Eletrônicos",
	// })

	db.Create(&Category{
		Name: "Cozinha",
	})

	// //create product
	// db.Create(&Product{
	// 	Name:       "Notebook",
	// 	Price:      1899.90,
	// 	CategoryID: 1,
	// })

	db.Create(&Product{
		Name:       "Panela",
		Price:      99.90,
		CategoryID: 1,
	})

	db.Create(&SerialNumber{
		Number:    "123456",
		ProductID: 1,
	})

	var categories []Category
	err = db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			fmt.Println("- ", product.Name, "Serial Number:", product.SerialNumber.Number)
		}
	}
}
