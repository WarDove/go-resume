package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(tpl.ParseFiles("kamran.gohtml", "tarlan.gohtml"))
}

func Resume(w http.ResponseWriter, r *http.Request) {
	if r.Host == "huseynov.net" || r.Host == "www.huseynov.net" {
		tpl.ExecuteTemplate(w, "tarlan.gohtml", nil)
	} else if r.Host == "kamran.huseynov.net" || r.Host == "www.kamran.huseynov.net" {
		tpl.ExecuteTemplate(w, "kamran.gohtml", nil)
	} else {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

}

func main() {
	fh := http.FileServer(http.Dir("./content"))
	http.HandleFunc("/", Resume)
	http.Handle("/css/", fh)
	http.Handle("/fonts/", fh)
	http.Handle("/gulzar/", fh)
	http.Handle("/img/", fh)
	http.Handle("/js/", fh)
	http.Handle("/scss/", fh)

	http.ListenAndServe(":80", nil)

}
