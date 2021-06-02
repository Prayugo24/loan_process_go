package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"tunaiku_tes/models"
	"github.com/joho/godotenv"
	"os"
	"fmt"
)

func SetupModels() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic("gagal Load env")
	}
	db_host := os.Getenv("DB_HOST")
	db_name := os.Getenv("DB_NAME")
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	config_db := db_user+":"+db_pass+"@("+db_host+":3306)/"+db_name+"?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql",config_db)
	
	fmt.Println(config_db)
	
	if err != nil {
		panic("gagal koneksi database")

	}
	db.AutoMigrate(&models.Tb_Invoice{},models.Tb_Loan{}, models.Tb_User_Loan{} )
	
	return db
}