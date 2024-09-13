package controller

import (
    "net/http"
    "github.com/RaihanMalay21/server-registry-tb-berkah-jaya-development/helper"
    config "github.com/RaihanMalay21/config-tb-berkah-jaya-development"
)

func LogOut(w http.ResponseWriter, r *http.Request) {
	// menghapus session 
	session, err := config.Store.Get(r, "berkah-jaya-session") 
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["id"] = ""
	session.Values["role"] = ""
	session.Options.MaxAge = -1 // expired session immediatly

	if err := session.Save(r, w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// menghapus cookie
	// Debug log
	http.SetCookie(w, &http.Cookie{
		Name : "token",
		Value: "",
		MaxAge: -1,
		Path:  "/",
		Secure: true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	})

	helper.Response(w, "Logout Berhasil", http.StatusOK)
	return
}