package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primary_key"`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	// db.Create(&Product{
	// 	Name:  "Produto 1",
	// 	Price: 10,
	// })

	// //create batch
	// products := []Product{
	// 	{Name: "Produto 2", Price: 20},
	// 	{Name: "Produto 3", Price: 30},
	// 	{Name: "Produto 4", Price: 40},
	// }
	// db.Create(&products)

	// var Product Product
	// db.First(&Product, 8)
	// fmt.Println(Product)
	// db.First(&Product, "name = ?", "Produto 4")
	// fmt.Println(Product)

	// var Products []Product
	// db.Find(&Products)
	// for _, product := range Products {
	// 	fmt.Println(product)
	// }

	// var Products []Product
	// db.Limit(2).Offset(2).Find(&Products)
	// for _, product := range Products {
	// 	fmt.Println(product)
	// }

	// var Products []Product
	// db.Where("price > ?", 20).Find(&Products)
	// for _, product := range Products {
	// 	fmt.Println(product)
	// }

	// var product Product
	// db.First(&product, 8)
	// product.Name = "Produto New"
	// db.Save(&product)

	// var product2 Product
	// db.First(&product2, 8)
	// fmt.Println(product2)
	// db.Delete(&product2)

}
