package main

import (
	"fmt"
)

func reverse(numbers []int, maxIndex int) []int {
	var reverseNumbers [10]int
	for i := 0; i < len(reverseNumbers); i++ {
		maxIndex--
		reverseNumbers[maxIndex] = numbers[i]
	}
	return reverseNumbers[:]
}

func main() {
	var numbers [10]int
	maxIndex := len(numbers)
	for i, _ := range numbers {
		fmt.Printf("Введите число %v: ", i+1)
		fmt.Scan(&numbers[i])
	}
	fmt.Println(reverse(numbers[:], maxIndex))
}
