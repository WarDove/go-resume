package main

import (
	"log"
	"net/http"
)

func redirectTLS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://"+r.Host+":443"+r.RequestURI, http.StatusMovedPermanently)

}
func main() {

	http.HandleFunc("/", redirectTLS)

	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatalln(err)
	}

}
