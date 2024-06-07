package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	tx := db.Begin()
	var category1 Category
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&category1, 1).Error
	if err != nil {
		panic(err)
	}

	category1.Name = "Eletrônicos"
	tx.Debug().Save(&category1)
	tx.Commit()

}
