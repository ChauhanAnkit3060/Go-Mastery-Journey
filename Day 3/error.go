package main

import "fmt"

func StartToPanic(a, b int) {
	if b == 0 {
		panic("World is ending, paaannnnniiiiicccccc")
	}

}

func Panicer() {
	StartToPanic(2, 0) // nothing will run after this
}

func Initiator() {
	done := make(chan bool)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered from Panic : ",r)
			}
			done <- true
		}()

		Panicer()
	}()

	<-done // Make recover blocking, until panic recovers or defer completes
	fmt.Println("Deeds endure !!!")
}