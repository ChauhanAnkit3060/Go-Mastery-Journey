package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func CheckPasswordStrength(pass string) (isValid bool, err error) {
	isValid = true
	err = nil
	if len(pass) < 8 {
		isValid = false
		err = fmt.Errorf("password must be at least 8 char long")
	}
	return
}

func main() {
	// Code to see who defer, panic and recover works
	Initiator()
	return
	fmt.Printf("Enter password : ")
	reader := bufio.NewReader(os.Stdin)
	pass, err := reader.ReadString('\n')
	if err != nil{
		log.Fatal(err)
		return
	}
	

	pass = strings.TrimSpace(pass)
	fmt.Printf("%v is of length %d\n",pass,len(pass))
	if isValid, err := CheckPasswordStrength(pass); !isValid{
		log.Fatal(err)
		return
	}
	fmt.Println("Password is valid")
}