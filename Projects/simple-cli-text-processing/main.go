package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"slices"
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
	data := make([]byte,10000)
	n,err := file.Read(data)
	if err == io.EOF{
		return string(data),nil
	}
	if err != nil{
		return "", fmt.Errorf("error while reading from file : %v",err)
	}
	return string(data[:n]),nil
}

func writeTofile(content, filepath string) error{
	file, err := os.Create(filepath)
	if err != nil{
		return fmt.Errorf("error while creating file %s : %v",filepath,err)
	}
	defer file.Close()
	n, err := file.Write([]byte(content))
	if err!=nil{
		return fmt.Errorf("error writing to file %s : %v",filepath,err)
	}
	fmt.Printf("%d byte written to file:%s successfully\n",n,filepath)
	return nil
}

func upperOrLower(str string, toUpper bool) string{
	if toUpper{
		return strings.ToUpper(str)
	}
	return strings.ToLower(str)
}

func strReverse(str string) string{
	strArr := strings.Split(str, " ")
	slices.Reverse(strArr)
	str = strings.Join(strArr," ")
	return str
}

func fileOption(readfile, writefile string, reader *bufio.Reader){
	content,err := readFromFile(readfile)
	if err!=nil{
		panic(err)
	}
	opt, _ := userInput("Please select operation (r: reverse date | u: all caps | l:lower case) : ",reader)
	switch opt{
	case "r" :
		content = strReverse(content)
		if err := writeTofile(content,writefile); err!=nil{
			panic(err)
		}
		fmt.Println("Contents of file are revsered")
	case "u" :
		content = upperOrLower(content,true)
		if err := writeTofile(content,writefile); err!=nil{
			panic(err)
		}
		fmt.Println("Contents of file are converted to uppercase")
	case "l" :
		content = upperOrLower(content,false)
		if err := writeTofile(content,writefile); err!=nil{
			panic(err)
		}
		fmt.Println("Contents of file are converted to uppercase")
	default :
		fmt.Println("Entered value is not valid ....")
		fileOption(readfile,writefile,reader)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var readfile, writefile string
	flag.StringVar(&readfile,"readfile","file.txt","File path to read from")
	flag.StringVar(&writefile,"writefile","output.txt","File path to write into")
	flag.Parse()

	content,_ := readFromFile(readfile)
	fmt.Printf("File: %s currently contains\n%s \n%s",readfile,strings.Repeat("-",50),content)
	fmt.Println()
	fileOption(readfile,writefile,reader)
}