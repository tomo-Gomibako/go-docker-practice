package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

func main() {
	count := 0
	db, err := sql.Open("mysql", "go:password@db/go?charset=utf8")
	// checkErr(err)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello!")
		// fmt.Fprintln(w, count)
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
