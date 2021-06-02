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


type V1ApproveLoanController struct {
	Status int
}

type inputApproveLoan struct {
	Status string
	Loan_date time.Time
	Time_period_loan time.Time
}

type LoanDataResult struct {
	Tenor int
	Status string
}

func (status *V1ApproveLoanController) ApproveLoan (c *gin.Context){
	db := c.MustGet("db").(*gorm.DB)
	IdUser, _ := strconv.ParseInt(c.PostForm("params[Id_user]"),0,64)
	DateApprove := c.DefaultPostForm("params[Date_approve]", "")

	var loansUserModels models.Tb_User_Loan

	if err := db.Where("id = ?", IdUser).First(&loansUserModels).Error; err != nil{
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "User Not Found",
		})
		return
	}
	
	LoanTb, err := db.Table("tb_user_loans as tsu").Select(" tsu.status, tbl.tenor ").Joins("JOIN tb_loans as tbl on tbl.id_user_loan = tsu.id").Where("tsu.id = ?",IdUser).Rows()
	defer LoanTb.Close()
	if err != nil {
        fmt.Println("error")
    }
	var LoanResult LoanDataResult
    for LoanTb.Next() {
        db.ScanRows(LoanTb, &LoanResult)
		fmt.Println(LoanResult)
    }
	if LoanResult.Tenor == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Tenor User not yet selected",
		})
		return
	}
	if LoanResult.Status == "Actived" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "User Is Approved",
		})
		return
	}
	dateApprove, err := time.Parse("2006-01-02",DateApprove);
	datePeriod := dateApprove.AddDate(0, 0, 29)
    if err !=nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Error while parsing date",
		})
		return
    }
	
	inputApprove := inputApproveLoan{
		Status : "Actived",
		Loan_date : dateApprove,
		Time_period_loan :datePeriod,
	}
	db.Model(&loansUserModels).Where("id = ?", IdUser).Update(inputApprove)
	
	readLoan, err := db.Table("tb_user_loans as tsu").Select("tsu.name , tsu.no_ktp, tsu.photo_ktp, tsu.date_of_birth, tsu.gender, tsu.address, tsu.current_job, tsu.no_hp, tsu.mothers_name, tsu.monthly_income, tsu.status, tbl.loan_interest_permonth, tbl.total_loan_fund, tbl.monthly_installments, tbl.tenor ").Joins("JOIN tb_loans as tbl on tbl.id_user_loan = tsu.id").Where("tsu.id = ?",IdUser).Rows()
	defer readLoan.Close()
	if err != nil {
        fmt.Println("error")
    }

	var result ResponseReadUserLoan
    for readLoan.Next() {
        db.ScanRows(readLoan, &result)
		fmt.Println(result)
    }

	c.JSON(200, gin.H{"status": 200, "response":result})
	return
}