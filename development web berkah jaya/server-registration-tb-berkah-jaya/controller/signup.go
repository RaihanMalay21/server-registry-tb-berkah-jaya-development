package controller

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"

    "golang.org/x/crypto/bcrypt"
    "github.com/go-playground/validator/v10"
    "github.com/go-playground/validator/v10/translations/id"

    "github.com/RaihanMalay21/server-registry-tb-berkah-jaya-development/helper"
    config "github.com/RaihanMalay21/config-tb-berkah-jaya-development"
    models "github.com/RaihanMalay21/models_TB_Berkah_Jaya"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	// mengambil json dari request body
	var UserSignup models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&UserSignup); err != nil {
		log.Println(err)
		fmt.Println("error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} 

	log.Println("UserSignup data:", UserSignup)

	// inalisasi validate 
	validate := validator.New(validator.WithRequiredStructEnabled())
	trans := helper.TranslatorIDN()
	id.RegisterDefaultTranslations(validate, trans)
	helper.RegisterCustomValidations(validate, trans)

	
	// Validasi data
	if err := validate.Struct(&UserSignup); err != nil {

		errs := err.(validator.ValidationErrors)
			
		errors := errs.Translate(trans)
		log.Println("Validation errors:", errors)
		helper.Response(w, errors, http.StatusBadRequest)
		return
	}

	// authentikasi apakah user nya sudah tersedia atau belum
	// var User models.User
	// if err := config.DB.Where("user_name = ?", UserSignup.UserName).First(&User).Error; err != nil{
	// 	switch err {
	// 	case gorm.ErrRecordNotFound:
			

	// 		// response 
	// 		helper.Response(w, "Berhasil Registrasi, Silahkan Login", http.StatusOK)
	// 		return
	// 	case nil:
	// 		// response jka use sudah ada
	// 		http.Error(w, "UserName Sudah Tersedia", http.StatusBadRequest)
	// 		return
	// 	default:
	// 		log.Println(err)
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// }

	// hash password menggunakan bcrypt
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(UserSignup.Password), bcrypt.DefaultCost)
	UserSignup.Password = string(hashPassword)

	// insert data user to database
	if err := config.DB.Create(&UserSignup).Error; err != nil {
		log.Println("Error Function Signup cant insert data:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	message := map[string]string{"message": "Berhasil Registrasi Silahkan Login"}
	helper.Response(w, message, http.StatusOK)
}