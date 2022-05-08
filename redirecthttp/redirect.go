package main

import (
	"log"
	"net/http"
)

func redirectTLS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://"+r.Host+r.RequestURI, http.StatusMovedPermanently)

}
func main() {

	if err := http.ListenAndServe(":80", http.HandlerFunc(redirectTLS)); err != nil {
		log.Fatalln(err)
	}

}
