package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// Page is the template for http response
type Page struct {
	Title string
	Count int
}

func main() {
	count := 0
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintln(w, "Hello, World!")
		page := Page{"Hello, World!", count}
		t, err := template.ParseFiles("index.html")
		if err != nil {
			panic(err)
		}
		err = t.Execute(w, page)
		if err != nil {
			panic(err)
		}
	})
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong")
	})
	http.HandleFunc("/countup", func(w http.ResponseWriter, r *http.Request) {
		count++
		fmt.Fprintln(w, "ok!")
	})
	http.ListenAndServe(":80", nil)
}
