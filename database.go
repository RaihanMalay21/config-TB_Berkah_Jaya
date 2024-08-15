package config 

import (
	"fmt"
	"os"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	models "github.com/RaihanMalay21/models_TB_Berkah_Jaya"
)

var (
	DB *gorm.DB
)

func DB_Connection() {
	// var (
	// 	dbUser = os.Getenv("DB_USER")
	// 	dbPwdd = os.Getenv("DB_PASSWORD")
	// 	dbHost = os.Getenv("DB_HOST")
	// 	dbPort = "3306"
	// 	dbName = os.Getenv("DB_NAME")
	// )

	// dbUser := "root"
	// dbPwdd := "0987"
	// dbHost := "localhost"
	// dbPort := "3306"
	// dbName := "TB_Berkah_Jaya"

	db, err := gorm.Open(mysql.Open("root:0987@tcp(localhost:3306)/TB_Berkah_Jaya?parseTime=true", dbUser, dbPwdd, dbHost, dbPort, dbName))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Barang{})
	db.AutoMigrate(&models.Hadiah{})
	db.AutoMigrate(&models.Pembelian{})
	db.AutoMigrate(&models.Pembelian_Per_Item{})
	db.AutoMigrate(&models.HadiahUser{})

	DB = db
}
