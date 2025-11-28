package main

import (
	"fmt"
	"os"
	"time"
)

type Bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func (b Bill) PrintBill() string {
	billStr := ""
	var billTotal float64
	billStr += fmt.Sprintln("Bill Name : ", b.name)
	billStr += fmt.Sprintf("Generated on : %s\n", time.Now().Format("02 Jan 06 15:04:05"))
	for item, val := range b.items {
		billStr += fmt.Sprintf("%s : %.2f\n",item,val)
		billTotal += val
	}
	billStr += fmt.Sprintf("CGST (2.5%%) : %.2f\n",0.025 * billTotal )
	billStr += fmt.Sprintf("SGST (2.5%%) : %.2f\n",0.025 * billTotal )
	billStr += fmt.Sprintf("Tip amount : %.2f\n",b.tip)
	billTotal = billTotal * 1.05 + b.tip
	billStr += fmt.Sprintf("Your total is : %.2f",billTotal)
	return billStr
}

func (b *Bill) UpdateTip(tip float64) {
	b.tip = tip
}

func (b *Bill) AddItem(item string, price float64) {
	b.items[item] = price
}

func (b Bill) SaveBill(){
	data := []byte(b.PrintBill())
	if err :=os.WriteFile("./bills/"+b.name+".txt",data,0644);err!=nil{
		panic(err)
	}
}