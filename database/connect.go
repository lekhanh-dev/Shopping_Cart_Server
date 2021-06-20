package database

import (
	"fmt"

	"github.com/lekhanh-dev/Shopping_Cart_Server/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/ecommerce?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("Connect database successful")
	}

	db.SetupJoinTable(&models.Order{}, "Products", &models.OrderDetail{})
	db.AutoMigrate(&models.Role{}, &models.User{}, &models.Category{}, &models.Product{}, &models.Order{})

}
