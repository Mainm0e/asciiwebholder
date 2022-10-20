package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/mainm0e/asciiweb/rary"
)

var temp *template.Template

func main() {

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/inputAscii", AsciiArtHandle)
	fmt.Println("Start sever at port 8080...")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func AsciiArtHandle(w http.ResponseWriter, r *http.Request) {
	type Data struct {
		Output    string
		ErrorNum  int
		ErrorText string
	}
	d := Data{}
	temp = template.Must(template.ParseGlob("static/*.html"))
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	if r.Method == "GET" {
		temp.ExecuteTemplate(w, "index.html", d)
	} else if r.Method == "POST" {
		Ascii := r.FormValue("inputText")
		Font := r.FormValue("Font")
		fmt.Println(rary.Output(Ascii, Font))
		d.Output = rary.Output(Ascii, Font)
	}
	temp.ExecuteTemplate(w, "index.html", d)

}
