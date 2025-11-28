package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
	io.Reader is capable for taking for input from
	-> os.File
	-> strings.Reader
	-> http.response.body
*/

func CountAlphabets(f io.Reader) (int, error){
	count := 0
	// buffer byte slice to store the data being read 
	data := make([]byte,1024)

	// read data into byte array
	n, err := f.Read(data)
	for _,c := range data[:n]{
		if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z'){
			count++
		}
	}
	if err == io.EOF{
		return count,err
	}
	if err != nil{
		return 0, err
	}
	return count,nil
}

/*
	io.Writer is capable for taking for input from
	-> os.File
	-> io.Stdio
	-> http.responseWritter
*/
func WrtieString(s string,  file io.Writer) (int, error){
	
	n,err := file.Write([]byte(s))
	if err!= nil{
		return 0, fmt.Errorf("error occured while writing error : %v",err)
	}
	return n,nil
}

func main() {
	// New reader to read from stdin
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter file path : ")
	filename, err := reader.ReadString('\n')
	filename = strings.TrimSpace(filename)
	if err != nil {
		panic(err)
	}

	// Open file to read
	file, err := os.Open(filename)
	if err != nil{
		panic(err)
	}
	// Close the file after use
	defer file.Close()

	count,err := CountAlphabets(file)
	if err != nil{
		panic(err)
	}
	fmt.Println("Letter count : ",count)

	// Read from string
	r := strings.NewReader("This is 💕")
	count, err = CountAlphabets(r)
	if err != nil{
		panic(err)
	}
	fmt.Println("Letter count : ",count)
	
	// Write into file
	fmt.Print("Enter file to write : ")
	filepath,_ := reader.ReadString('\n')
	filepath = strings.TrimSpace(filepath)
	file1, err := os.Create(filepath)
	if err!=nil{
		panic(err)
	}
	defer file1.Close()

	WrtieString(" 💕 This is newly created file\n This is first line in it \nOverwrite 💕",file1)

	file2,_ := os.OpenFile("file.txt",os.O_APPEND,0644)
	defer file2.Close()
	file2.Write([]byte("\n this is a new log entry\n"))
	data,_ := os.ReadFile("file.txt")
	fmt.Println("Data : ",string(data))

}