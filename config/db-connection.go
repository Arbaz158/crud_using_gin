package config

import (
	"fmt"

	"github.com/crud_using_gin/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DbConnection() *gorm.DB {
	url := "root:rootwdp@tcp(localhost:3306)/demo"
	Db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		fmt.Println("error in getting connection :", err)
	}
	Db.AutoMigrate(&entity.Video{})
	return Db
}
