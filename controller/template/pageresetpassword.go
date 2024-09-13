package template

import (
    "html/template"
    "log"
    "net/http"
)

func PageResetPassword(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("resetPassword.html")
	if err != nil {
		log.Println("Error Cant parse Files html:", err.Error())
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println("Error cant execute html tamplate:", err.Error)
		return
	}
}

