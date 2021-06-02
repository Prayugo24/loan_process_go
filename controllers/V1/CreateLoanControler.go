package V1

import(
	"github.com/gin-gonic/gin"
    "tunaiku_tes/models"
	"tunaiku_tes/helpers"
	"fmt"
	"net/http"
	"github.com/google/uuid"
	"path/filepath"
	"github.com/jinzhu/gorm"
	"strconv"
	"math"
	"time"
	"strings"
	age "github.com/bearbin/go-age"

	
)
type LoanUserData struct {
	ID   int `json:"id"`
    Name string `bson:"name"` 
	No_ktp int `bson:"no_ktp"`
}
type Tenors struct {
	Month_3 float64
	Month_6 float64
	Month_12 float64
}
type ResponseCreateUserLoan struct {
	Id int
	Name string
	No_ktp int64
	Photo_ktp string
	Date_of_birth string
	Gender string
	Address string
	No_hp int
	Mothers_name string
	Status string
	Total_loan_fund float64
	Loan_interest_permonth float64
	Tenor Tenors
	
}

type V1CreateLoanController struct {
	Status int
}

func (status *V1CreateLoanController) CreateLoan (c *gin.Context){
	db := c.MustGet("db").(*gorm.DB)
	NoKtp, _ := strconv.ParseInt(c.PostForm("params[No_ktp]"),0,64)
	Name := c.DefaultPostForm("params[Name]", "")
	PlaceOfBirth := c.DefaultPostForm("params[Place_of_birth]", "")
	DateOfBirth := c.DefaultPostForm("params[Date_of_birth]", "")
	Gender := c.DefaultPostForm("params[Gender]", "")
	Address := c.DefaultPostForm("params[Address]", "")
	CurrentJob := c.DefaultPostForm("params[Current_job]", "")
	NoHp, _ := strconv.Atoi(c.PostForm("params[No_hp]"))
	MothersName := c.DefaultPostForm("params[Mothers_name]", "")
	MonthlyIncome, _ := strconv.ParseFloat(c.PostForm("params[Monthly_income]"),64)
	Provinces := strings.ToLower(c.DefaultPostForm("params[Province]", ""))

	parasmLoanValidation := helpers.ValidationIncomeParams{
		MonthlyIncome: MonthlyIncome,
	}
	LoanValidation := helpers.ValidationIncome(parasmLoanValidation)
	if LoanValidation.Status != 200{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "your income does not meet the minimum requirements of 3 million",
		})
		return
	}
	layout := "2006-01-02"
	dates, _ := time.Parse(layout, DateOfBirth)
	ageUser := age.Age(dates)
	if ageUser < 17 || ageUser > 80{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Your age does not meet",
		})
		return
	}
	
	if Provinces != "dki jakarta" && Provinces != "jawa barat" && 
		Provinces != "jawa timur" && Provinces != "sumatera utara"{
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "The province is not listed",
			})
			return
	}
	
	LoanUsers, _ := db.Table("tb_user_loans").Where("no_ktp = ?",NoKtp).Rows()
	defer LoanUsers.Close()
    // Values to load into
    result := make([]LoanUserData, 0)
    for LoanUsers.Next() {
        var row LoanUserData
        result = append(result, row)
    }
	if len(result) > 0 {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "The data already exists",
		})
		return
	}
	// proces Photo_ktp_face_photo
	PhotoKtpFacePhoto, err := c.FormFile("params[Photo_ktp_face_photo]")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}
	// Retrieve file information
	extensions := filepath.Ext(PhotoKtpFacePhoto.Filename)
	// Generate random file name for the new uploaded file so it doesn't override the old file with same name
	PhotoKtpFacePhotoName := uuid.New().String() + extensions
	// The file is received, so let's save it
	if err := c.SaveUploadedFile(PhotoKtpFacePhoto, "assets/images/" + PhotoKtpFacePhotoName); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}
	
	 // proces Photo_ktp
	 PhotoKtp, err := c.FormFile("params[Photo_ktp]")
	 if err != nil {
		 c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			 "message": "No file is received Photo Ktp and Foto Face",
		 })
		 return
	 }
	 // Retrieve file information
	 var extension string = filepath.Ext(PhotoKtp.Filename)
	 // Generate random file name for the new uploaded file so it doesn't override the old file with same name
	 PhotoKtpName := uuid.New().String() + extension
	 // The file is received, so let's save it
	 
	 if err := c.SaveUploadedFile(PhotoKtp, "assets/images/" + PhotoKtpName); err != nil {
		 c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			 "message": "Unable to save the file Photo Ktp",
		 })
		 return
	 }
	dateLoan := time.Now()
	datePeriod := dateLoan.AddDate(0, 0, 29)
	inputLoanUser := models.Tb_User_Loan{
		No_ktp : NoKtp,
		Name : Name,
		Photo_ktp : PhotoKtpName,
		Photo_ktp_face_photo : PhotoKtpFacePhotoName,
		Place_of_birth : PlaceOfBirth,
		Date_of_birth : DateOfBirth,
		Gender : Gender,
		Address : Address,
		Current_job : CurrentJob,
		No_hp : NoHp,
		Mothers_name : MothersName,
		Monthly_income : MonthlyIncome,
		Province:Provinces,
		Status:"Deactived",
		Loan_date: dateLoan,
		Time_period_loan :datePeriod,
	}
	
	db.Create(&inputLoanUser)
	
	inputLoanModels := models.Tb_Loan{
		Handling_fee : LoanValidation.HandlingFee,
		Loan_interest_permonth : LoanValidation.LoanInterestPermonth,
		Late_fee : LoanValidation.LateFee,
		Total_loan_fund : LoanValidation.LoanUser,
		Tenor : 0,
		Monthly_installments : 0,
		Id_user_loan : inputLoanUser.Id,
	}
	fmt.Println("Name  :",inputLoanModels)
	db.Create(&inputLoanModels)
	
	response := ResponseCreateUserLoan{
		Id : inputLoanUser.Id,
		Name : inputLoanUser.Name,
		No_ktp : inputLoanUser.No_ktp,
		Photo_ktp : inputLoanUser.Photo_ktp,
		Date_of_birth : inputLoanUser.Date_of_birth,
		Gender : inputLoanUser.Gender,
		Address : inputLoanUser.Address,
		No_hp : inputLoanUser.No_hp,
		Mothers_name : inputLoanUser.Mothers_name,
		Status : inputLoanUser.Status,
		Total_loan_fund: inputLoanModels.Total_loan_fund,
		Loan_interest_permonth: inputLoanModels.Loan_interest_permonth,
		Tenor:Tenors{
			Month_3 : math.Ceil(((inputLoanModels.Total_loan_fund/3)*100)/100),
			Month_6 : math.Ceil(((inputLoanModels.Total_loan_fund/6)*100)/100),
			Month_12 : math.Ceil(((inputLoanModels.Total_loan_fund/12)*100)/100),
		},
	}
    if LoanValidation.Status == 200 {
		c.JSON(200, gin.H{"status": 200,"response":response})
    }else{
		c.JSON(400, gin.H{"status": 400, "response":""})
	}
	return
	
}


