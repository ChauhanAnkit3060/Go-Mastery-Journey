package main

import "fmt"

type Bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func (b Bill) PrintBill() {
	var billTotal float64
	fmt.Println("Bill Name : ", b.name)
	for item, val := range b.items {
		fmt.Printf("%s : %.2f\n",item,val)
		billTotal += val
	}
	fmt.Printf("CGST (2.5%%) : %.2f\n",0.025 * billTotal )
	fmt.Printf("SGST (2.5%%) : %.2f\n",0.025 * billTotal )
	fmt.Printf("Tip amount : %.2f\n",b.tip)
	billTotal = billTotal * 1.05 + b.tip
	fmt.Printf("Your total is : %.2f",billTotal)
}

func (b *Bill) UpdateTip(tip float64) {
	b.tip = tip
}

func (b *Bill) AddItem(item string, price float64) {
	b.items[item] = price
}