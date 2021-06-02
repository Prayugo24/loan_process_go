package V1

import(
	"github.com/gin-gonic/gin"
    "tunaiku_tes/models"
	"fmt"
	"github.com/jinzhu/gorm"
	"strconv"
	"net/http"
	"math"
	
)

type V1CreateTenorController struct {
	Status int
}

type LoanInput struct {
	Tenor int
	Monthly_installments float64
	Loans_received float64
}

type LoanData struct {
	Total_loan_fund float64
	Handling_fee float64
	Loans_received float64
	Loan_interest_permonth float64
	Tenor int
	Monthly_installments float64
	Id_user_loan int
	Id int
}

func (status *V1CreateTenorController) CreateTenor (c *gin.Context){
	db := c.MustGet("db").(*gorm.DB)
	IdUser, _ := strconv.ParseInt(c.PostForm("params[Id_user]"),0,64)
	Tenors, _ := strconv.Atoi(c.PostForm("params[Tenor]"))

	var loansModels models.Tb_Loan

	if err := db.Where("id_user_loan = ?", IdUser).First(&loansModels).Error; err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	LoanTb, err := db.Table("tb_loans").Where("id_user_loan = ?",IdUser).Rows()
	defer LoanTb.Close()
	if err != nil {
        fmt.Println("error")
    }
    // Values to load into
	var loanData LoanData
    for LoanTb.Next() {
		db.ScanRows(LoanTb, &loanData)
        fmt.Println(loanData)
    }
	
	var MonthlyInstallments float64
	if Tenors == 3 {
		MonthlyInstallments = math.Ceil(((loanData.Total_loan_fund/3)*100)/100)
	}else if Tenors == 6 {
		MonthlyInstallments = math.Ceil(((loanData.Total_loan_fund/6)*100)/100)
	}else if Tenors == 12 {
		MonthlyInstallments = math.Ceil(((loanData.Total_loan_fund/12)*100)/100)
	}else{
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "The tenor must be filled in as 3 or 6 or 12 months",
		})
		return
	}
	totalMonthlyInstallments := (MonthlyInstallments+loanData.Loan_interest_permonth)
	LoansReceived := loanData.Total_loan_fund - loanData.Handling_fee
	inputLoan := LoanInput{
		Tenor : Tenors,
		Monthly_installments : totalMonthlyInstallments,
		Loans_received : LoansReceived ,
	}
	
	db.Model(&loansModels).Update(inputLoan)
	response := LoanData{
		Id : loanData.Id,
		Handling_fee : loanData.Handling_fee,
		Loan_interest_permonth : loanData.Loan_interest_permonth,
		Total_loan_fund : loanData.Total_loan_fund,
		Loans_received : LoansReceived,
		Tenor : Tenors,
		Monthly_installments : totalMonthlyInstallments,
		Id_user_loan : loanData.Id_user_loan,
	}
	if response.Id != 0 {
		c.JSON(200, gin.H{"status": 200,"response":response})
    }else{
		c.JSON(400, gin.H{"status": 400, "response":""})
	}
	return
}