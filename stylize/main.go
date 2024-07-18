package main

import (
	"fmt"
	"html/template"
	"net/http"

	"Fonctions/Fonctions"
)

type ascii struct {
	Result, Banner, Text string
}

var q ascii

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("POST /", ascii_Art)
	fmt.Println("Server is Runing...")
	http.ListenAndServe(":8012", nil)
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
	temp.Execute(w, nil)
}

func ascii_Art(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Status Bad Request 400", http.StatusBadRequest)
		return
	}
	q.Text = r.Form.Get("text")
	if q.Text == "" {
		q.Text = "type something..."
	}
	q.Banner = r.Form.Get("banner")
	if q.Banner != "standard.txt" && q.Banner != "shadow.txt" && q.Banner != "thinkertoy.txt" {
		http.Error(w, "Status Bad Request 400", http.StatusBadRequest)
		return
	}
	q.Result = Fonctions.AsciiArt(q.Banner, q.Text)
	if q.Result == "" {
		http.Error(w, "Internal Server Error: status 500", http.StatusBadRequest)
		return
	}
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
