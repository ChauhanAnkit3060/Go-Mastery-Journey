package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Hello Git...")
	fmt.Printf("Time now : %v",time.Now())

	var dia_txt string	
	fmt.Print("\nEnter diameter of circle : ")	
	fmt.Scanln(&dia_txt) 
	fmt.Println("Entered diameter : ",dia_txt)
	dia , err := strconv.ParseFloat(dia_txt,64)
	if err != nil{
		fmt.Println("Invalid response")
		return
	}
	fmt.Printf("Circumference : %.2f\n",calcCircumference(dia))
	// The '&' before 'name' passes the memory address of the variable	
	// fmt.Printf("Hello, %s!\n", name)	var age int	
	// fmt.Print("Enter your age: ")	fmt.Scanln(&age)	
	// fmt.Printf("You are %d years old.\n", age)}

	reader := bufio.NewReader(os.Stdin) // Create a new reader for standard input	
	fmt.Print("Enter a sentence: ")	
	sentence, _ := reader.ReadString('\n') // Read until a newline character	
	fmt.Printf("You entered: %s", sentence) // sentence will include the newline character}
}

//git config --global user.name "ChauhanAnkit3060"
//git config --global user.email "theankitchauhan3060@gmail.com"

