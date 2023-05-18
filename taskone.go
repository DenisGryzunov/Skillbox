package main

import (
	"fmt"
)

func main() {
	var numbers [10]int
	for i, _ := range numbers {
		fmt.Printf("Введите число %v: ", i+1)
		fmt.Scan(&numbers[i])
	}
	var counterEvenNumber int
	var counterOddNumber int

	for _, number := range numbers {
		if number%2 == 0 {
			counterEvenNumber++
		} else {
			counterOddNumber++
		}
	}
	fmt.Printf("Количество четных чисел: %d\nКоличество нечетных чисел: %d", counterEvenNumber, counterOddNumber)
}
