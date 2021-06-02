package models

import( 
	"time"
)

type Tb_Invoice struct {
	Id int `json:"id", gorm:"primary_key", gorm:autoIncrement` // id
	Date time.Time `json:"date"`
	Monthly_bil float64 `json:"monthly_bil"`
	Amercement float64 `json:"amercement"`
	Id_user_loan int64 `json:"id_user_loan"`
}