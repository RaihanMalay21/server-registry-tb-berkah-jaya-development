package middlewares

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	helper "github.com/RaihanMalay21/helper_TB_Berkah_Jaya"
	"golang.org/x/net/http2"
)

// function ReverseProxy is forward request user to target microserveces
func ReverseProxy(target string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url, err := url.Parse(target)
		if err != nil {
			log.Println("Error Parsing URL: ", err)
			message := map[string]interface{}{"message": err}
			helper.Response(w, message, http.StatusInternalServerError)
			return
		}

		// create a new reverse proxy to the target microservices
		Proxy := httputil.NewSingleHostReverseProxy(url)

		// set custom Transport To support HTTP/s
		Proxy.Transport = &http2.Transport{}

		// modify request
		r.URL.Host = url.Host
		r.URL.Scheme = url.Scheme
		r.Header.Set("X-Forwarded-Host", r.Host)
		r.Host = url.Host

		Proxy.ServeHTTP(w, r)
	})
}