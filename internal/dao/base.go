package dao

import (
	"caipiaotong/internal/models"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func InitDB() {
	var err error
	addr := viper.GetString("mysql.addr")
	user := viper.GetString("mysql.user")
	password := viper.GetString("mysql.password")
	dbName := viper.GetString("mysql.dbName")
	charset := viper.GetString("mysql.charset")
	dsn := user + ":" + password + "@tcp(" + addr + ")/" + dbName + "?charset=" + charset

	log.Println("dsn:", dsn)
	db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Println(err)
		return
	}
	err = db.AutoMigrate(
		&models.User{},
		&models.Bill{},
	)
	if err != nil {
		panic(err)
		return
	}
	log.Println("database connected")
}
