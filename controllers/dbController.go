package controllers

import (
	"log"

	"github.com/kakty/taskmgr/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//InitDb :
func InitDb() *gorm.DB {
	db, err := gorm.Open(postgres.Open("user=kakty dbname=todos port=5432 sslmode=disable "), &gorm.Config{})
	if err != nil {
		log.Fatal("Error opening db")
	}
	err := db.Table("tasks").AutoMigrate(&models.Task{})
        if err != nil{
                log.Fatal(err)
        }
	//db.Create(&models.Task{ID: 1, Name: "nam", Description: "des", StartTime: time.Now(), FinishTime: time.Now().Add(time.Second * 3600), IsCompleted: false})
	err := db.Table("users").AutoMigrate(&models.User{})
        if err != nil {
                log.Fatal(err)
        }
	db.Table("users").Create(&models.User{Username: "kakty", Password: "3574"})
	return db
}
