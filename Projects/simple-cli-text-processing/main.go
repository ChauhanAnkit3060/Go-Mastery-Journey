package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)



func main() {
	var readfile, writefile string
	flag.StringVar(&readfile,"readfile","readfile.txt","file to read from !")
	flag.StringVar(&writefile,"writefile","filewrite.txt","file to write in !")
	flag.Parse()
	fmt.Println(readfile,writefile)

	data, err := os.ReadFile(readfile)
	if err!=nil{
		panic(err)
	}
	strdata := string(data)
	strdata = strings.ReplaceAll(strdata,"\n","\n😈")
	if err :=os.WriteFile(writefile,[]byte(strdata),0644);err!=nil{
		panic(err)
	}

	fmt.Println("Program completed")
	
}