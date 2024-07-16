package main

import (
	"fmt"
	"html/template"
	"net/http"

	"Fonctions/Fonctions"
)

type ascii struct {
	Result string
	Banner string
	Text   string
}

var q ascii

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("POST /", ascii_Art)
	fmt.Println("Server is Runing...")
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Status Not Found 404", http.StatusNotFound)
		return
	}
	temp, err := template.ParseFiles("src/index.html")
	if err != nil {
		http.Error(w, "Status Internal Server Error 500", http.StatusInternalServerError)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Status Bad Request 400", http.StatusBadRequest)
		return
	}
	q.Text = r.Form.Get("text")
	q.Banner = r.Form.Get("banner")
	fmt.Println(q.Text, q.Banner)
	if q.Banner == "" || q.Text == "" {
		temp.Execute(w, nil)
		return
	}
	q.Result = Fonctions.PrintAsciiArt(q.Banner, q.Text)
	temp.Execute(w, q)
}

func ascii_Art(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Status Method Not Allowed 404", http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Status Bad Request 400", http.StatusBadRequest)
		return
	}
	str := r.Form.Get("text")
	banner := r.Form.Get("banner")
	res := Fonctions.AsciiArt(banner, str)
	var q ascii
	q.Result = res
	q.Banner = banner
	if len(str) > 1 && str[0:2] == "\r\n" {
		q.Text = "\r\n" + str
	} else {
		q.Text = str
	}
	temp, err := template.ParseFiles("src/index.html")
	if err != nil {
		http.Error(w, "Status Internal Server Error 500", http.StatusInternalServerError)
		return
	}
	temp.Execute(w, q)
}
