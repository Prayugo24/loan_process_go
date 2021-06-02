package routers

import (
	"github.com/gin-gonic/gin"
	"tunaiku_tes/config"
	"net/http"
	"tunaiku_tes/controllers/V1"
	"github.com/aviddiviner/gin-limit"
)

func RouterMain() http.Handler  {
	router := gin.New()
	db := config.SetupModels()

	router.Use(func(c * gin.Context){
		c.Set("db",db)
		c.Next()
	})
	router.Use(limit.MaxAllowed(50))
	v1CreateLoan := &V1.V1CreateLoanController{Status: 200}
	v1CreateTenor := &V1.V1CreateTenorController{Status: 200}
	v1ReadLoan := &V1.V1ReadLoanController{Status: 200}
	v1UpdateLoan := &V1.V1UpdateLoanController{Status: 200}
	v1ApproveLoan := &V1.V1ApproveLoanController{Status: 200}

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{"data":"Welcome To Api Golang"})
	})

	router.POST("/create_loan", v1CreateLoan.CreateLoan)
	router.POST("/create_tenor", v1CreateTenor.CreateTenor)
	router.POST("/read_loan", v1ReadLoan.ReadLoan)
	router.POST("/update_loan", v1UpdateLoan.UpdateLoan)
	router.POST("/approve_loan", v1ApproveLoan.ApproveLoan)

	

	return router
}
