//Implementation of common sorting algorithms for practice
package sorts

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

//create randomized slice of ints
type rList []int

//Newlist is a constructor for rList.
func NewList(size int) rList {
	rl := make(rList, size)
	rl.Randomize()
	return rl
}

//Randomize inserts new random values into list
func (rl rList) Randomize() {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < len(rl); i++ {
		rl[i] = random.Intn(math.MaxInt64)
	}
}

//Reverse the ordering of elements.
func (rl rList) Reverse() {
	tmp := make(rList, len(rl))
	for i := 0; i < len(rl); i++ {
		tmp[i] = rl[len(rl)-i-1]
	}
	for i := 0; i < len(rl); i++ {
		rl[i] = tmp[i]
	}
}

//Sorted checks if the slice of ints is sorted
//in either increasing or decreasing order.
//if empty, or all equal values, then returns true
func Sorted(list []int) bool {
	if sortedInc(list) == false {
		fmt.Println("sortedInc returned false...")
		if sortedDec(list) == false {
			fmt.Println("sortedDec returned false...")
			return false
		}
	}
	return true
}

func sortedInc(list []int) bool {
	for i := 0; i < len(list)-1; i++ {
		if list[i] > list[i+1] {
			fmt.Println("%v > %v?", list[i], list[i+1])
			return false
		}
		if i == len(list)-2 {
			return true
		}
	}
	return true
}

func sortedDec(list []int) bool {
	for i := 0; i < len(list)-1; i++ {
		if list[i] < list[i+1] {
			return false
		}
		if i == len(list)-2 {
			return true
		}
	}
	return true
}

//Insertion sort has n^2 complexity.
func Insertion(slice []int) {
	for j := 1; j < len(slice); j++ {
		key := slice[j]
		i := j - 1
		for i >= 0 && slice[i] > key {
			slice[i+1] = slice[i]
			i--
		}
		slice[i+1] = key
	}
}

//Merge sort has nlg(n) complexity.
func Merge(slice []int) {
	mergeSort(slice, 0, len(slice)-1)
}

//mergeSort divides the list recursively into sub-sections by calling mergeSort(),
//and calls 'merge()' on each sub-section (p..r) to sort.
func mergeSort(slice []int, p, r int) { //p is first element, r is last element
	if p < r { // if more than 1 element in slice, otherwise already sorted
		q := int(math.Floor(float64((p + r) / 2))) // middle element used to split slice
		mergeSort(slice, p, q)
		mergeSort(slice, q+1, r)
		merge(slice, p, q, r)
	}
}

//called on a sub-section of list, where each half is sorted (p..q) and (q+1..r),
//and put into separate slices, which are used to compare and sort entire sub-section (p..r)
func merge(slice []int, p, q, r int) {
	var A, B []int
	//loops fill up slices with appropriate values from slice
	for i := p; i <= q; i++ {
		A = append(A, slice[i])
	}
	for i := q + 1; i <= r; i++ {
		B = append(B, slice[i])
	}
	A = append(A, math.MaxInt64) //add sentinel value to end of each slice
	B = append(B, math.MaxInt64)
	i, j := 0, 0
	//A and B are sorted, so loop compares and adds lowest of each slice
	//to p..r in slice
	for k := 0; k <= r-p; k++ {
		if A[i] <= B[j] {
			slice[p+k] = A[i]
			i++
		} else {
			slice[p+k] = B[j]
			j++
		}
	}
}
