package main

import (
	"fmt"
	"os"
)

func LoadConfig(filename string) (string, error) {
	if filename == "" {
		return "", fmt.Errorf("config path cannot be empty")
	}
	filedata,err := os.ReadFile(filename)
	if err!=nil{
		return "", fmt.Errorf("error : %v",err)
	}
	return string(filedata),nil
}

func main() {
	// mux := http.NewServeMux()
	// mux.HandleFunc("/v1/Hello",func(w http.ResponseWriter,r *http.Request){
	// 	fmt.Fprintf(w,"Hello Go")
	// })

	// http.ListenAndServe(":8080",mux)

	filename := ""
	str,err := LoadConfig(filename)
	if err!=nil{
		fmt.Fprintf(os.Stderr,"Error : %v\n",err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout,"File content : \n\n %s",str)
}