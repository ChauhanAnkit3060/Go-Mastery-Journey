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

	// Slices
	// Long declaration
	var cores []int
	fmt.Println(cores,len(cores))
	cores = append(cores,101)
	fmt.Println(cores,len(cores))
	
	newGames := games[1:4]
	fmt.Println(newGames,len(newGames))

	// maps
	fmt.Printf("\nBlock for Maps Demo\n")
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
		fmt.Printf("%s\t\t:\t\t%2.f\n",key,val)
	}
}