package middlewares

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"golang.org/x/net/http2"
	"github.com/RaihanMalay21/api-gateway-tb-berkah-jaya-development/helper"
)


func ReverseProxy(target string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		message := map[string]interface{} {
			"message": nil,
		}

		url , err := url.Parse(target)
		if err != nil {
			log.Println(err)
			message["message"] = err.Error()
			helper.Response(w, message, http.StatusInternalServerError)
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(url)

		// Set Custom Transport To Support HTTP/2
		proxy.Transport = &http2.Transport{}

		pathMicroservices := []string{
			"http://localhost:8081",
			"http://localhost:8082",
			"http://localhost:8083",
		}

		pathPrefix := []string{
			"/customer",
			"/access",
			"/admin",
		}

		for i, path := range pathMicroservices {
			if path == url.String() {
				r.URL.Path = strings.TrimPrefix(r.URL.Path, pathPrefix[i])
				break
			}
		}


		// set the request to forward into services
		r.URL.Host = url.Host
		r.URL.Scheme = url.Scheme
		r.Header.Set("X-Forwarded-Host", r.Host)
		r.Host = url.Host

		proxy.ServeHTTP(w, r)
	})
}