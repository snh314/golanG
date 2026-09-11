package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
	// Declare a package-level variable 'db'
	// Type: *gorm.DB (pointer to GORM database connection)
	// This stores the database connection for entire application
)

func Connect() {
	d, err := gorm.Open("mysql", "root:@/bookstore?charset=utf8&parseTime=True&loc=Local")
	// gorm.Open() opens database connection
	// Parameter 1: "mysql" = database driver type
	// Parameter 2: connection string = "user:password@/dbname?options..."
	//    user     = database username
	//    password = database password
	//    dbname   = database name
	if err != nil {
		panic(err)
		// Check if connection failed, panic stops the program immediately
	}
	db = d
}
func GetDB() *gorm.DB {
	return db
}
