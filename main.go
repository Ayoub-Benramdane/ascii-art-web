package main

import (
	"html/template"
	"io"
	"net/http"
)

var s = `
<!DOCTYPE html>
<html>
<body>

<h2>HTML Links</h2>
<p>HTML links are defined with the a tag:</p>
<h1>{{.}}</h1>
<a href="/test">This is a link</a>

</body>
</html>
`
type test struct{
	name string
	lname string
}

func index(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte(s))

	if r.URL.Path == "/" {
		st := ""
		r.ParseForm()
		st += r.Form.Get("fname")
		sv := r.Form.Get("lname")
		
		t := test{st,sv}
		temp := template.Must(template.ParseFiles("./stc/index.html")) 
		
		temp.Execute(w, t)
		// st += r.PostFormValue("fname")
		// st += r.PostFormValue("lname")

		// io.WriteString(w, s)
		return
	}
	io.WriteString(w, "error 404 ")
}

func main() {
	http.HandleFunc("/", index)
	k := http.FileServer(http.Dir("./stc/"))
	http.Handle("/about/", http.StripPrefix("/about/", k))
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./stc/help.html")
	})
	http.ListenAndServe(":8089", nil)
}
