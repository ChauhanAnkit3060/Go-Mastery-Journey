package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func userInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	if err != nil{
		return "",fmt.Errorf("error while taking input from stdio : %v",err)
	}
	input = strings.TrimSpace(input)
	return input, nil
}

func readFromFile(filename string) (string, error){
	file,err := os.Open(filename)
	if err != nil{
		return "", fmt.Errorf("error while opening from file : %v",err)
	}
	defer file.Close()
	data := make([]byte,1000)
	n,err := file.Read(data)
	if err == io.EOF{
		return string(data),nil
	}
	if err != nil{
		return "", fmt.Errorf("error while reading from file : %v",err)
	}
	return string(data[:n]),nil
}

func filePrompt(){
	fmt.Print("Please select file operator (n - create) ")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	name, _ := userInput("Enter file : ",reader)
	fmt.Print(name)
}