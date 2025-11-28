package main

import "fmt"

func fibonacci(n int, ch chan int, quit chan bool) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	quit<-true
}

func main() {
	ch := make(chan int)
	quit := make(chan bool)
	go fibonacci(10, ch, quit)
	loop :
	for{
		select {
		case i := <-ch:
			fmt.Printf("%d | ",i)
		case <-quit :
			fmt.Println()
			fmt.Println("Program ended !")
			break loop
		}
	}
}