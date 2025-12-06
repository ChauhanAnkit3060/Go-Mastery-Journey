package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct{
	Id int `json:"userId"`
	Name string `json:"userName"`
	Email string `json:"userEmail"`
}

func handleHome(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Welcome to home !!")
}

func handleTime(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Time now is : %s",time.Now().Format(time.RFC1123))
}

func handleUser(w http.ResponseWriter, r *http.Request){
	fmt.Println("URL :",r.URL.Path)
	fmt.Println("Qery for name : ",r.URL.Query().Get("name"))
	User := User{
		Name: "Ankit",
		Id : 12,
		Email: "test@test.com",
	}

	data,err := json.MarshalIndent(User,"","\t")
	if err!= nil{
		http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
		return
	}
	data = append(data, '\n')
	fmt.Println(string(data))
	fmt.Fprintf(w,"user : %s",string(data))

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}