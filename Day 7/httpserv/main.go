package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/",handleHome)
	mux.HandleFunc("/time",handleTime)
	mux.HandleFunc("/api/user",handleUser)

	http.ListenAndServe(":8080",mux)
}