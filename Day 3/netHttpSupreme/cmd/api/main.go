package main

import (
	"log"
	"net/http"
)

func main() {
	// 1. Go is smart enough to resolve conflict for similar path, by using the most specific one
	// 		But if path are equally specific then program panics
	router := http.NewServeMux()
	router.HandleFunc("/caseid/{id}",func (w http.ResponseWriter, r * http.Request)  {
		id := r.PathValue("id")
		w.Write([]byte("Received request for case id "+ id))
	})
	router.HandleFunc("/caseid/lastest",func (w http.ResponseWriter, r * http.Request)  {
		w.Write([]byte("This is Latest"))
	})
	// 2. While associating the path to Mux we can specify the HTTP methods
	//    This will cause all other request on same path go "Method not allowed"
	router.HandleFunc("POST /caseid/Create/{id}",func (w http.ResponseWriter, r * http.Request)  {
		id := r.PathValue("id")
		w.Write([]byte("Created ID : " + id))
	})

	server := http.Server{
		Handler: router,
		Addr: ":8080",
	}

	log.Println("Starting server on port :8080")
	server.ListenAndServe()
}