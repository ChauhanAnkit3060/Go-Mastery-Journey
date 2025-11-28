package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func UserInput(prompt string, r *bufio.Reader)(string,error){
	fmt.Print(prompt)
	name,err := r.ReadString('\n')
	name = strings.TrimSpace(name)
	return name,err
}

func CreateBill() Bill {
	reader := bufio.NewReader(os.Stdin)
	name,_ := UserInput("Enter bill name : ",reader)
	bill := Bill{name: name,items: make(map[string]float64)}

	promptAll(bill)
	return bill
}

func promptAll(b Bill){
	reader := bufio.NewReader(os.Stdin)
	opt,_ := UserInput("Please select an option (a : add item | s : save bill | t : add tip) : ",reader)
	switch opt{
	case "a" :
		item,_ := UserInput("Add item name : ",reader)
		p,_ := UserInput("Enter item cost : ",reader)
		price,err := strconv.ParseFloat(p,64)
		if err!=nil || price <0{
			fmt.Println("Price can only be a number...")
			promptAll(b)
		}
		b.AddItem(item,price)
		fmt.Println("Item added to bill")
		promptAll(b)
	case "s" :
		fmt.Print("Saving bill....\n")
		b.SaveBill()
		fmt.Print("Bill Saved Successfully !!\n")
	case "t" :
		t,_ := UserInput("Enter tip amount : ",reader)
		tip,err := strconv.ParseFloat(t,64)
		if err!=nil{
			fmt.Println("Tip can only be a number...")
			promptAll(b)
		}
		b.UpdateTip(tip)
		fmt.Println("Tip added to bill, thanks a lot !!")
		promptAll(b)
	default :
		fmt.Println("selected option is not correct .....")
		promptAll(b)
	}
}

func main() {
	/*
	bill := Bill{
		name: "Office Lunch",
		items: map[string]float64{
			"Veg Combo":     85,
			"Egg Combo":     90,
			"Chineis Combo": 87,
			"Dessert":       60,
			"Cold drinks":   110,
		},
		tip: 20,
	}
	*/
	bill := CreateBill()
	fmt.Println(bill)
}