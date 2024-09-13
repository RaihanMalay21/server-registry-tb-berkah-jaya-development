package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"github.com/RaihanMalay21/server-registry-tb-berkah-jaya-development/helper"
	models "github.com/RaihanMalay21/models_TB_Berkah_Jaya"
	config "github.com/RaihanMalay21/config-tb-berkah-jaya-development"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// mengambil request json
	var Userlogin map[string]string
	JSON := json.NewDecoder(r.Body)
	if err := JSON.Decode(&Userlogin); err != nil {
		log.Println("Error Decode JSON:", err)
		http.Error(w, "Gagal login!, Silahkan coba lagi", http.StatusInternalServerError)
		return
	}
	
	// inialisasi session
	session, err := config.Store.Get(r, "berkah-jaya-session")
		if err != nil {
			log.Println("Error Getting session:", err)
			http.Error(w, "cannot sign to session", http.StatusInternalServerError)
			return
		}

	// jika yang login adalah admin 
	if Userlogin["usernameORemail"] == "RaihanMalay21" || Userlogin["usernameORemail"] == "Wirawati21" || Userlogin["usernameORemail"] =="Yondrizal21" {
		// mengambil data dari database
		var adminlogin models.User
		if err := config.DB.Where("user_name = ?", Userlogin["usernameORemail"]).First(&adminlogin).Error; err != nil {
			switch err {
			case gorm.ErrRecordNotFound:
				log.Println("Admin username tidak ditemukan:", err)
				msg := map[string]string{"message": "Username Tidak di Temukan"}
				helper.Response(w, msg, http.StatusBadRequest)
				return
			default:
				log.Println("Error query admin user:", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		// komperasi password antara admin dan di db
		if err := bcrypt.CompareHashAndPassword([]byte(adminlogin.Password), []byte(Userlogin["password"])); err != nil {
			switch err {
			case bcrypt.ErrMismatchedHashAndPassword:
				log.Println("Password mismatch:", err)
				msg := map[string]string{"message": "Password Salah"}
				helper.Response(w, msg, http.StatusBadRequest)
				return
			default:
				log.Println("Error comparing password:", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		// membuat token
		expTime := time.Now().Add(24 * time.Hour)
		claims := &config.JWTClaim{
			UserName: Userlogin["usernameORemail"],
			ID: adminlogin.ID, 
			Role: "Admin",
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer: "go-jwt-mux",
				ExpiresAt: jwt.NewNumericDate(expTime),
			},
		}

		// mendeklarasi algoritma yang akan digunakan untuk signed token
		tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		
		// signed toke
		token, err := tokenAlgo.SignedString(config.JWT_KEY)
		if err != nil {
			log.Println("Error Signing token:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// set token ke dalam cookie
		http.SetCookie(w, &http.Cookie{
			Name: "token",
			Value: token,
			HttpOnly: true,
			Secure: true, 
			Path: "/",
			MaxAge: 24 * 60 * 60,
			// SameSite: http.SameSiteLaxMode,
			SameSite: http.SameSiteNoneMode, // mengizinkan lintas domain
		})

		// set session untuk menyimpan data sensitif users 
		session.Values["id"] = adminlogin.ID
		session.Values["role"] = "Admin"

		if err := session.Save(r, w); err != nil {
			log.Println("Error saving session:", err)
			http.Error(w, "cannot save session", http.StatusInternalServerError)
			return
		}

		helper.Response(w, "Login Berhasil", http.StatusOK)
		return
	}


	// auhentikasi apakah email 
	var fieldColumn string
	usernameORemail, ok := Userlogin["usernameORemail"]
	if ok && usernameORemail != "" {

		if strings.Contains(usernameORemail, "@") {
			fieldColumn = "email"
		} else {
			fieldColumn = "user_name"
		}

	} 

	if fieldColumn == "" {
		http.Error(w, "Username atau email harus diisi", http.StatusBadRequest)
		return
	}

	// mengambil data berdarkan username
	var login models.User
	if err := config.DB.Where(fieldColumn + " = ?", usernameORemail).First(&login).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			log.Println("User username tidak ditemukan:", err)
			msg := map[string]string{"message": "Username Tidak di Temukan"}
			helper.Response(w, msg, http.StatusBadRequest)
			return
		default:
			log.Println("Error query user:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// komperasi password antara usercustomers dan customers
	if err := bcrypt.CompareHashAndPassword([]byte(login.Password), []byte(Userlogin["password"])); err != nil {
		switch err {
		case bcrypt.ErrMismatchedHashAndPassword:
			log.Println("Password mismatch:", err)
			msg := map[string]string{"message": "Password Salah"}
			helper.Response(w, msg, http.StatusBadRequest)
			return
		default:
			log.Println("Error comparing password:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// proses pembuatan token jwt
	expTime := time.Now().Add(24 * time.Hour)
	claims := &config.JWTClaim{
		UserName: login.UserName,
		ID: login.ID,
		Role: "Customers",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "go-jwt-mux",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	// mendeklarasi algoritma yang akan digunakan untuk signed token
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// signed token
	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// set token ke dalam cookie
	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: token, 
		HttpOnly: true,
		Secure: true,
		Path: "/",
		MaxAge: 24 * 60 * 60,
		// SameSite: http.SameSiteLaxMode,
		SameSite: http.SameSiteNoneMode,
	})

	// menyimpan data sensitif ke dalam session
	session.Values["id"] = login.ID
	session.Values["role"] = "Customers"

	if err := session.Save(r, w); err != nil {
		log.Println("Error saving session:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helper.Response(w, "Login Berhasil", http.StatusOK)
	return
}