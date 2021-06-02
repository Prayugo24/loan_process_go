package models

import( 
	"time"
)

type Tb_User_Loan struct {
	Id int `json:"id", gorm:"primary_key", gorm:autoIncrement` // id
	Name string `json:"name"` 
	No_ktp int64 `json:"no_ktp"` 
	Photo_ktp string `json:"photo_ktp"` 
	Photo_ktp_face_photo string `json:"photo_ktp_face_photo"` 
	Place_of_birth string `json:"place_of_birth"` 
	Date_of_birth string `json:"date_of_birth"` 
	Gender string `json:"gender"` 
	Address string `json:"address"` 
	Current_job string `json:"current_job"` 
	No_hp int `json:"no_hp"` 
	Mothers_name string `json:"mothers_name"` 
	Monthly_income float64 `json:"monthly_income"` 	
	Province string `json:"province"` 
	Status string `json:"status"` 
	Loan_date time.Time `json:"loan_Date"` 
	Time_period_loan time.Time `json:"time_period_loan"` 


}