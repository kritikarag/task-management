package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func ConnectDatabase() (*gorm.DB, error) {
	user := "root"
	pass := "Wanderlust@26"
	host := "localhost"
	port := "3306"
	dbname := "task_management"
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, port, dbname)
	db, err := gorm.Open("mysql", URL) 
	if err != nil {
		panic(err.Error())
	}
	DB = db
	// Optional: If you want to see the SQL logs
	// db.LogMode(true)

	return db,nil
}


// func ConnectDatabase() *gorm.DB{
// 	USER := "kritikarag"
// 	PASS := "password"
// 	HOST := "localhost"
// 	PORT := "8084"
// 	DBNAME := "task_management"
// 	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)

// 	db, err := gorm.Open("mysql", URL)

// 	if err!=nil{
// 		panic("Failed to connect to database!")
// 	}
// 	//database.AutoMigrate(&Task{})
// 	DB = db
// 	return db
// }
