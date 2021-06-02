package V1

import(
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/jinzhu/gorm"
	"tunaiku_tes/models"
	"strconv"
	"time"
	"net/http"
	
)


type V1UpdateLoanController struct {
	Status int
}

type LoanUserRespons struct {
	Status string
	Loan_date string
	Monthly_installments float64
	Late_fee float64
	Time_period_loan time.Time
}

func (status *V1UpdateLoanController) UpdateLoan (c *gin.Context){
	db := c.MustGet("db").(*gorm.DB)
	IdUser, _ := strconv.ParseInt(c.PostForm("params[Id_user]"),0,64)
	PayLoan, _ := strconv.ParseFloat(c.PostForm("params[Pay_loan]"),64)
	DatePayLoan := c.DefaultPostForm("params[Date_pay_loan]", "")

	LoanUsers, _ := db.Table("tb_user_loans as tsu").Select("tsu.status, tsu.loan_Date, tsu.time_period_loan, tbl.monthly_installments, tbl.late_fee").Joins("JOIN tb_loans as tbl on tbl.id_user_loan = tsu.id").Where("tsu.id = ?",IdUser).Rows()
	defer LoanUsers.Close()
	var responLoans LoanUserRespons
    for LoanUsers.Next() {
        db.ScanRows(LoanUsers, &responLoans)
		fmt.Println(responLoans)
    }
	if responLoans.Status != "Actived"{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Loan application has not been approved",
		})
		return
	}

	if PayLoan != responLoans.Monthly_installments{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "The amount of payment must be the same as the monthly bill",
		})
		return
	}
	datePay, err := time.Parse("2006-01-02",DatePayLoan);
	if err !=nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Error while parsing date",
		})
		return
    }
	cekDatePay := datePay.Before(responLoans.Time_period_loan)
	var amercement float64
	var mounthBil float64
	if cekDatePay {
		// cek if date less then Time_period_loan
		mounthBil = PayLoan
		amercement = 0
	}else{
		daysLate := datePay.Sub(responLoans.Time_period_loan).Hours() / 24
		amercement = responLoans.Late_fee * daysLate
		mounthBil = PayLoan + amercement
	}
    
	inputInvoiceLoan := models.Tb_Invoice{
		Date : datePay,
		Monthly_bil : mounthBil,
		Id_user_loan : IdUser,
		Amercement : amercement,
	}
	db.Create(&inputInvoiceLoan)

	c.JSON(200, gin.H{"status": 200,"response":inputInvoiceLoan})
	return
}