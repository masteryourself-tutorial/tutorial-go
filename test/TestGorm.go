package main

import (
	"chat/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dns := "root:123456@tcp(127.0.0.1:3306)/chat?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.User{})

	// Create
	db.Create(&model.User{Name: "张三"})

	// Read
	var user model.User
	db.First(&user, 1)             // find user with integer primary key
	db.First(&user, "id = ?", "1") // find user with code D42

	// Update - update user's price to 200
	db.Model(&user).Update("password", "123456")
	// Update - update multiple fields
	db.Model(&user).Updates(model.User{Name: "李四"}) // non-zero fields
	db.Model(&user).Updates(map[string]interface{}{"email": "lisi@qq.com", "password": "123"})

	// Delete - delete user
	db.Delete(&user, 1)
}
