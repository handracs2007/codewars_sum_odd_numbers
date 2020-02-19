package main

import "fmt"

func RowSumOddNumbers(n int) int {
	return n * n * n
}

func main() {
	fmt.Println(RowSumOddNumbers(1))
	fmt.Println(RowSumOddNumbers(2))
	fmt.Println(RowSumOddNumbers(13))
}
