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
	// http.HandleFunc("POST /", ascii_Art)
	http.HandleFunc("/ascii-art", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	})

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	// 1
	// file, err := os.ReadFile("src/index.html")
	// if err != nil {
	// 	// w.WriteHeader(http.StatusInternalServerError)
	// w.Write([]byte("ghhfg"))
	// 	return
	// }

	// fmt.Fprint(w, string(file))

	// 2
	// http.ServeFile(w, r, "src/index.html")

	// 3
	// http.Handle("/",http.StripPrefix("/",http.FileServer(http.Dir("./src/"))))

	// 4
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
	// 1
	// str := r.FormValue("text")
	// w.Write([]byte(str))
	// 2

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
	res := Fonctions.PrintAsciiArt(banner, str)
	var q ascii
	q.Result = res
	q.Banner = banner
	q.Text = str
	temp, err := template.ParseFiles("src/index.html")
	if err != nil {
		http.Error(w, "Status Internal Server Error 500", http.StatusInternalServerError)
		return
	}
	temp.Execute(w, q)
}
