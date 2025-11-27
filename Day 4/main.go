package main

import "fmt"

func main() {
	// Arrays
	//	are of fixed size
	//	homoginious collection

	// Long declaration
	// All value will be blank untill initalize
	var names [10]string = [10]string{"ddew"}
	fmt.Println(names)
	names[5] = "Ankit"
	fmt.Println(names)
	
	// Short declaration
	games := [10]string{"GOW","40K","LOZ","ER","DS"}
	fmt.Println(games)
	for _,s := range games{
		fmt.Printf("%s | ",s)
	}
	println()

	fmt.Printf("\nBlock for Slice Demo\n\n")
	// Slices
	// Long declaration
	var cores []int
	fmt.Println(cores,len(cores))
	cores = append(cores,101)
	fmt.Println(cores,len(cores))
	
	newGames := games[1:4]
	fmt.Println(newGames,len(newGames))

	newSlice := make([]int,0,10)
	fmt.Println(len(newSlice))

	// maps
	fmt.Printf("\nBlock for Maps Demo\n\n")
	menus := map[string]float64{
		"Tea" : 10.00,
		"Onion Baji" : 25.00,
		"Veg Cutlet" : 30.00,
		"Maggie": 35.00,
	}
	fmt.Println(menus)
	menus["Coffee"] = 15.00

	fmt.Printf("Fixed Menus\n")
	for key,val := range menus{
		fmt.Printf("%s : %2.f\n",key,val)
	}

	item := "Tea"
	if v,ok := menus[item]; ok{
		fmt.Printf("%s is available at cost %2.f",item,v)
	}else{
		fmt.Printf("%s is not available",item)
	}

	// Interface
	fmt.Printf("\n\n This block is for Interface code \n\n")
	shapes := []Shape{
		Sqaure{4.3},
		Circle{6.43},
		Circle{93.4},
		Sqaure{89.3},
	}
	for _,sq := range shapes{
		PrintInfo(sq)
	}

	//Exercise 
	nums := []int{43,2,235,25,23,2352,1241,436,7457,4,3,214,2}
	fmt.Println("Exercise answer :",SumEvenNumber(nums))
}