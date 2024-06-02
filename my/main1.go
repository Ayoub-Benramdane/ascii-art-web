package main

import (
	"Fonctions/Fonctions"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("GET /", index)
	http.HandleFunc("POST /", ascii_Art)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	// 1

	// file, err := os.ReadFile("src/index.html")
	// if err != nil {
	// 	http.Error(w,"Status Internal Server Error 500",http.StatusInternalServerError)
	// 	return
	// }

	// fmt.Fprint(w, string(file))

	// 2

	// http.ServeFile(w, r, "src/index.html")

	// 3

	// http.Handle("/",http.StripPrefix("/",http.FileServer(http.Dir("./src/"))))

	// 4

	temp, err := template.ParseFiles("src/index.html")
	if err != nil {

		// w.WriteHeader(http.StatusInternalServerError)
		// w.Write([]byte("ghhfg"))

		http.Error(w, "Status Internal Server Error 500", http.StatusInternalServerError)

		return
	}
	temp.Execute(w, nil)
}

func ascii_Art(w http.ResponseWriter, r *http.Request) {
	// 1

	// str := r.FormValue("text")

	// 2

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Status Bad Request 400", http.StatusBadRequest)
		return
	}
	str := r.Form.Get("text")
	banner := r.Form.Get("banner")
	res := Fonctions.PrintAsciiArt(banner, str)
	w.Write([]byte(res))
	// w.Write([]byte(banner))
}
