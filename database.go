package config

import (
	"fmt"

	models "github.com/RaihanMalay21/models_TB_Berkah_Jaya"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func DB_Connection() {
	var (
		dbUser string
		dbPwdd string
		dbHost string
		dbName string
	)

	dbUser = getParameter("DB_USER")
	dbPwdd = getParameter("DB_PASSWORD")
	dbHost = getParameter("DB_HOST")
	dbName = getParameter("DB_NAME")

	dbURI := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", dbUser, dbPwdd, dbHost, dbName)

	db, err := gorm.Open(mysql.Open(dbURI))
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
