package main

import "fmt"

func updateByValue(a, b int) {
	a += 10
	b += 5
}

func updateByRef(a, b *int) {
	*a += 10
	*b += 5
}

func main() {

	counter := Counter{}
	fmt.Println(counter)
	counter.Increment()
	counter.Increment()
	counter.Increment()
	fmt.Println(counter)



	var (
		a int
		b int
	)
	fmt.Println("A : ",a," B : ",b)
	updateByValue(a,b)
	fmt.Println("A : ",a," B : ",b)
	updateByRef(&a,&b)
	fmt.Println("A : ",a," B : ",b)

	var user1, user2 User
	user1.InitializeUser("Kartik")
	fmt.Println("User : ",user1)
	user2.InitializeUser1("Piyush")
	fmt.Println("User : ",user2)

	userA := User{}

	// 1. Using the Pointer Receiver method
	userA.InitializeUser("Alice")
	userA.Deactivate()
	fmt.Println(userA) // Output: {1001 Alice true} - The original struct is updated.

	// 2. Reset and use the Value Receiver method
	userB := User{}
	userB.InitializeUser1("Bob")
	fmt.Println(userB) // Output: {0  false} - The original struct is NOT updated.
}