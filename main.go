package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(tpl.ParseFiles("kamran.gohtml", "tarlan.gohtml"))
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}

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
func getInstance() string {
	resp, err := http.Get("http://169.254.169.254/latest/meta-data/instance-id")
	checkErr(err)
	bs := make([]byte, resp.ContentLength)
	resp.Body.Read(bs)

	return string(bs)

}

func Instance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	io.WriteString(w, getInstance())
}

func main() {
	fh := http.FileServer(http.Dir("./content"))
	http.HandleFunc("/", Resume)
	http.HandleFunc("/instance", Instance)
	http.Handle("/css/", fh)
	http.Handle("/fonts/", fh)
	http.Handle("/gulzar/", fh)
	http.Handle("/img/", fh)
	http.Handle("/js/", fh)
	http.Handle("/scss/", fh)

	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatalln(err)
	}

}
