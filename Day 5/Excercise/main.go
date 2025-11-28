package main

func sliceSum(nums []int, ch chan int) {
	sum := 0
	for _,i := range nums {
		sum += i
	}
	ch <- sum
}

func main(){
	ch := make(chan int)
	nums := []int{4,2,5,3,23,5,3,2,33,2}
	go sliceSum(nums[:len(nums)/2],ch)
	go sliceSum(nums[len(nums)/2:],ch)
	x,y := <-ch,<-ch
	println(x,y,x+y)
}