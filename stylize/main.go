package main

import (
	"fmt"
	"log"
	"net/http"

	"Fonctions/server"
)

func main() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css/"))))
	http.HandleFunc("/", server.Index)
	http.HandleFunc("/ascii-art", server.Ascii_Art)
	fmt.Println("Server is Runing...")
	fmt.Println("Link : http://localhost:8404/")
	err := http.ListenAndServe(":8404", nil)
	if err != nil {
		log.Fatal(err)
	}
}
