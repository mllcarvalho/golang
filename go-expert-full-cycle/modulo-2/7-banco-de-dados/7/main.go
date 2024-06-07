package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID         int `gorm:"primary_key"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:product_categories;"`
	gorm.Model
}

type Category struct {
	ID       int `gorm:"primary_key"`
	Name     string
	Products []Product `gorm:"many2many:product_categories;"`
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

	category1 := Category{
		Name: "Eletrônicos",
	}
	db.Create(&category1)

	category2 := Category{
		Name: "Cozinha",
	}
	db.Create(&category2)

	db.Create(&Product{
		Name:       "Torradeira",
		Price:      1000.00,
		Categories: []Category{category1, category2},
	})

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			println("- ", product.Name)
		}
	}
}
