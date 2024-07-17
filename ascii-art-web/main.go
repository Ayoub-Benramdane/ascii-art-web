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
	http.ListenAndServe(":8404", nil)
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
	temp.Execute(w, q)
}

func ascii_Art(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Status Method Not Allowed 405", http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Status Bad Request 400", http.StatusBadRequest)
		return
	}
	q.Text= r.Form.Get("text")
	q.Banner = r.Form.Get("banner")
	if q.Banner != "standard.txt" && q.Banner != "shadow.txt"  && q.Banner != "thinkertoy.txt" {
		http.Error(w, "Status Bad Request 400", http.StatusBadRequest)
		return
	}
	q.Result = Fonctions.AsciiArt(q.Banner, q.Text)
	if len(q.Text) > 1 && q.Text[0:2] == "\r\n" {
		q.Text = "\r\n" + q.Text
	}
	temp, err := template.ParseFiles("src/index.html")
	if err != nil {
		http.Error(w, "Status Internal Server Error 500", http.StatusInternalServerError)
		return
	}
	temp.Execute(w, q)
}
