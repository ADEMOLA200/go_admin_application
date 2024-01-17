package database

import (
	"github.com/ADEMOLA200/go_admin_application/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(mysql.Open("root:rootroot@/go_admin_application"), &gorm.Config{})

	if err != nil {
		panic("cannot connect to database")
	}
	
	// db.Migrator().DropTable(&models.User{})
	// db.Migrator().CreateTable(&models.User{})

	DB = db

	DB.AutoMigrate(&models.User{})
}