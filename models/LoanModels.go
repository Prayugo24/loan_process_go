package models

type Tb_Loan struct {
	Id int `json:"id", gorm:"primary_key", gorm:autoIncrement` // id
	Handling_fee float64 `json:"handling_fee"` // biaya penagnanan biaya jasa
	Loan_interest_permonth float64 `json:"loan_interest_permonth"` //bunga perbulan 
	Late_fee float64 `json:"late_fee"` // biaya keterlambatan pembayaran
	Total_loan_fund float64 `json:"loan_fund"` // total dana pinjaman yang di terima
	Tenor int `json:"tenor"` // jangka waktu pemabayaran dalam bulanan
	Monthly_installments float64 `json:"monthly_installments"` //angsuran perbulan
	Loans_received float64 `json:"loans_received"` // pinjaman yang di terima setelah dikurangi biaya admin
	Id_user_loan int `json:"id_user_loan"`
}