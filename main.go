package main

import (
	"fmt"
	"github.com/tendingnut/Algorithms/sorts"
)

func main() {
	fmt.Println("this is main...")
	slice := []float64{3, 6, 2, 3.56, 23, 3, 3, 1, 2}
	sorted := sorts.Insertion(slice)
	reverse := sorts.Reverse(sorted)
	fmt.Println(slice, sorted, reverse)
}
