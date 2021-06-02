package helpers

import (
	"math"
)

type ValidationIncomeParams struct {
    MonthlyIncome float64
}

type ResponValidate struct{
	Status int
	LoanUser float64
	HandlingFee float64
	LoanInterestPermonth float64
	LateFee float64
}


func ValidationIncome(params ValidationIncomeParams) (response ResponValidate){
    if params.MonthlyIncome >= 1000000 || params.MonthlyIncome <=10000000 {
		loanUser := (params.MonthlyIncome/3)
		// handlingFee := (loanUser * (0.05))
		loanInterestPermonth := ((0.025)* loanUser)
		lateFee := loanUser*(0.05)
		response = ResponValidate{
			Status : 200,
			LoanUser : math.Ceil((loanUser*100)/100),
			HandlingFee : 100000,
			LoanInterestPermonth :math.Ceil((loanInterestPermonth*100)/100),
			LateFee : math.Ceil((lateFee*100)/100),
		}
	}else{
		response = ResponValidate{
			Status : 400,
			LoanUser : 0,
			HandlingFee : 0,
			LoanInterestPermonth : 0,
			LateFee : 0,
		}
	}
	return response
}