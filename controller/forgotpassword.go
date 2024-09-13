package controller

import (
    "fmt"
    "log"
    "net/http"
    "strings"

    "gorm.io/gorm"

    "github.com/RaihanMalay21/server-registry-tb-berkah-jaya-development/helper"
    config "github.com/RaihanMalay21/config-tb-berkah-jaya-development"
    models "github.com/RaihanMalay21/models_TB_Berkah_Jaya"
    middlewares "github.com/RaihanMalay21/middlewares_TB_Berkah_Jaya"
)

func ForgotPassword(w http.ResponseWriter, r *http.Request) {
	emailOrUsername := r.FormValue("emailOrUsername")

	fmt.Println(emailOrUsername)

	var fieldColumn string
	var errorMessageType string
	if emailOrUsername != "" {
		
		if strings.Contains(emailOrUsername, "@") {
		   fieldColumn = "email"	
		   errorMessageType = "Email Tidak Di Temukan" 
		} else {
		   fieldColumn = "user_name"
		   errorMessageType = "Username Tidak Di Temukan"
		}

	} else {
	   log.Println("Error emailOrUsername empty in function forgotPassword")
	   message := map[string]string{"message": "username or email kosong"}
	   helper.Response(w, message, http.StatusBadRequest)
	   return
	}

	// check whether email or username exist or not
	var user models.User
	if err := config.DB.Where(fieldColumn + "= ?", emailOrUsername).First(&user).Error; err != nil {
	   switch err{
	   case gorm.ErrRecordNotFound:
		   log.Println("Error data not found on function ForgotPassword:", err)
		   message := map[string]interface{}{"message": errorMessageType}
		   helper.Response(w, message, http.StatusBadRequest)
		   return
	   default:
		   log.Println("Error cant retreaving data from database function ForgotPassword:", err)
		   http.Error(w, err.Error(), http.StatusInternalServerError)
		   return
	   }
	}

   // generate a reset token
   tokenstr, err := middlewares.GenerateResetToken(user.Email)
   if err != nil {
	   log.Println("Error cannot generate a reset token function ForgotPassword:", err)
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
   } 

   // send email with token before have create
   if err := helper.SendEmail(user.Email, user.UserName, tokenstr, "ForgotPassword", ""); err != nil {
	   log.Println("Error cant send Email to user function ForgotPassword:", err)
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
   }

   var message = map[string]string{"message": "Silahkan Tunggu kami Mengirim Gmail"}
   helper.Response(w, message, http.StatusOK)
}