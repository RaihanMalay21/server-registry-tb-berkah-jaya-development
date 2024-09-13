package controller

import (
    "html/template"
    "log"
    "net/http"

    "golang.org/x/crypto/bcrypt"
    "github.com/golang-jwt/jwt/v5"
    config "github.com/RaihanMalay21/config-tb-berkah-jaya-development"
    models "github.com/RaihanMalay21/models_TB_Berkah_Jaya"
    middlewares "github.com/RaihanMalay21/middlewares_TB_Berkah_Jaya"
)

func ForgotPasswordChangePassword(w http.ResponseWriter, r *http.Request) {
	password := r.FormValue("password")
	passwordRepeat := r.FormValue("passwordRepeat")
	token := r.FormValue("token")

	type DataMessage struct {
		Password string
		PasswordRepeat string
		NotMatched string
		Token string
		Error interface{}
	}

	dataMessage := DataMessage{
		Password: password, 
		PasswordRepeat: passwordRepeat, 
		Token: token,
	}

	tmpl, err := template.ParseFiles("././template/resetPassword.html")
	if err != nil {
		log.Println("Error cant parse file template html :", err.Error())
		return
	}

	email, err := middlewares.VerifyResetToken(token)
	if err != nil {
		switch err {
		case jwt.ErrTokenSignatureInvalid:
			log.Println("Error Token signature invalid function ForgotPasswordChangePassword:", err)
			dataMessage.Error = "Token Tidak Valid"
		case jwt.ErrTokenExpired:
			log.Println("Error token has expired:", err)
			dataMessage.Error = "Token Telah Habis Waktunya"
		default:
			log.Println("Error Cant verify Token In Function ForgotPasswordChangePassword:", err)
			dataMessage.Error = err.Error()
		}

		if err := tmpl.Execute(w, dataMessage); err != nil {
			log.Println("Error cant execute template html:", err.Error())
			return
		}
		return
	}

	if password != passwordRepeat {
		log.Println("Password not match bettwen password and passwordRepeat funtion forgotPasswordChangePassword")
		dataMessage.NotMatched = "Password Tidak Sesuai"
		dataMessage.PasswordRepeat = ""

		if err := tmpl.Execute(w, dataMessage); err != nil {
			log.Println("Error cant execute template html:", err.Error())
			return
		}
		return
	}

	// hashing password 
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hashPasswordString := string(hashPassword)


	// update in database on field column password users
	if err := config.DB.Model(&models.User{}).Where("email = ?", email).Update("password", hashPasswordString).Error; err != nil {
		log.Println("Error Cant updated password in database function ForgotPasswordChangePassword:", err.Error())
		dataMessage.Error = "Error Cant Update Passowrd"

		if err := tmpl.Execute(w, dataMessage); err != nil {
			log.Println("Error cant execute template html:", err.Error())
			return
		}
		return
	}

	message := map[string]string{"message": "Berhasil Reset Password, Silahkan Login"}
	if err := tmpl.Execute(w, message); err != nil {
		log.Println("Error Cant execute template html:", err.Error())
	}
}