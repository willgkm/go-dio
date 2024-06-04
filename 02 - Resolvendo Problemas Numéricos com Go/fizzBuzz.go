package main

import "fmt"

func main() {
	var i int8 = 1
	for ; i <= 100; i++ {
		if dividedByThreeAndFive(i) {
			printBuzzFizz() 
		} else if dividedByThree(i) {
			printFizz()
		} else if dividedByFive(i) {
			printBuzz()
		} else if !dividedByThreeAndFive(i) {
			fmt.Println(i)
		}
	}
}

func dividedByThree(number int8) bool {
	return number%3 == 0
}
func dividedByFive(number int8) bool {
	return number%5 == 0
}

func dividedByThreeAndFive(number int8 ) bool {
	return dividedByThree(number) && dividedByFive(number)
} 

func printFizz() {
	fmt.Println("Fizz")
}

func printBuzz() {
	fmt.Println("Buzz")
}

func printBuzzFizz() {
	fmt.Println("BuzzFizz")
}
