package main

func SumEvenNumber(nums []int) int {
	var sum int
	for _,val := range nums{
		if val%2==0{
			sum+=val
		}
	}
	return sum
}