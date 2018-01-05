package main

import (
	"fmt"
	"github.com/tendingnut/Algorithms/sorts"
)

var (
	list = []float64{24, 53, 123, 3, 2, 5, 324, 5.3, 53, 5.3, 5.35, 5.2, 53.2, 523, 9}
)

func main() {
	fmt.Printf("Starting main...\n")
	sorts.Merge(list)
	fmt.Println(list)
}
