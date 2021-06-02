package V1

import(
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/jinzhu/gorm"
	"strconv"
	
)


type V1ReadLoanController struct {
	Status int
}
type ResponseReadUserLoan struct {
	Name string
	No_ktp int64
	Photo_ktp string
	Date_of_birth string
	Gender string
	Address string
	Current_job string 
	No_hp int
	Mothers_name string
	Monthly_income float64 
	Status string
	Loan_interest_permonth float64
	Total_loan_fund float64
	Monthly_installments float64 
	Loans_received float64
	Tenor int
}

func (status *V1ReadLoanController) ReadLoan (c *gin.Context){
	db := c.MustGet("db").(*gorm.DB)
	IdUser, _ := strconv.ParseInt(c.PostForm("params[Id_user]"),0,64)

	readLoan, err := db.Table("tb_user_loans as tsu").Select("tsu.name , tsu.no_ktp, tsu.photo_ktp, tsu.date_of_birth, tsu.gender, tsu.address, tsu.current_job, tsu.no_hp, tsu.mothers_name, tsu.monthly_income, tsu.status, tbl.loan_interest_permonth, tbl.total_loan_fund, tbl.monthly_installments, tbl.tenor, tbl.loans_received ").Joins("JOIN tb_loans as tbl on tbl.id_user_loan = tsu.id").Where("tsu.id = ?",IdUser).Rows()
	defer readLoan.Close()
	if err != nil {
        fmt.Println("error")
    }

	var result ResponseReadUserLoan
    for readLoan.Next() {
        db.ScanRows(readLoan, &result)
		fmt.Println(result)
    }

	if err == nil {
		c.JSON(200, gin.H{"status": 200,"response":result})
    }else{
		c.JSON(400, gin.H{"status": 400, "response":""})
	}
	return
}