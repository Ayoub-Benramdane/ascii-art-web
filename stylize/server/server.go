package server

import (
	"html/template"
	"net/http"

	"Fonctions/Fonctions"
)

type ascii struct {
	Result, Banner, Text string
}
var q ascii


func Index(w http.ResponseWriter, r *http.Request) {
	
	if r.URL.Path != "/" {   // ERROR 404(/) ***
		http.Error(w, "Status Not Found 404", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {    // ERROR 405 ***
		http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
		return
	}

	temp, err := template.ParseFiles("src/index.html")
	if err != nil {       // ERROR 500 (ParseFiles(...html))  
		http.Error(w, "Status Internal Server Error 500", http.StatusInternalServerError)
		return
	}

	err1 := temp.Execute(w, nil)     
	if err1 != nil {    // ERROR 500 (Execute(w, nil)) ***
		return
	}
}


func Ascii_Art(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {   // Bad --> internal ***
		http.Error(w, "Status Interal server 500", http.StatusBadRequest)
		return
	}

	if r.URL.Path != "/ascii-art" {       // ERROR 404(/ascii-art) ***
		http.Error(w, "Page NOT found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodPost {    // ERROR 405 ***
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	q.Text = r.Form.Get("text")        // capture the text
	if q.Text == "" {
		q.Text = "type something..."
	}

	q.Banner = r.Form.Get("banner")   // capture the banner
	if q.Banner != "standard.txt" && q.Banner != "shadow.txt" && q.Banner != "thinkertoy.txt" {    // ERROR 400
		http.Error(w, "Status Bad Request 400", http.StatusBadRequest)
		return
	}

	q.Result = Fonctions.AsciiArt(q.Banner, q.Text)
	if q.Result == "" {              // ERROR 500
		http.Error(w, "Internal Server Error: status 500", http.StatusBadRequest)
		return
	}

	if len(q.Text) > 1 && q.Text[0:2] == "\r\n" {
		q.Text = "\r\n" + q.Text
	}

	temp, err := template.ParseFiles("src/index.html")
	if err != nil {       // ERROR 500 (ParseFiles(...html)) 
		http.Error(w, "Status Internal Server Error 500", http.StatusInternalServerError)
		return
	}
	err1 := temp.Execute(w, q)
	if err1 != nil {    // ERROR 500 (Execute(w, q)) ***
		return
	}
}
