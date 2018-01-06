//Implementation of common sorting algorithms for practice, without
//using package sort from the standard library
package sorts

import (
	"math"
	"math/rand"
	"time"
)

//Used to easily create randomized lists for sorting.
type list []int

//Newlist is a constructor for list.
func NewList(size, maxInt int) list {
	l := make(list, size)
	l.Randomize(maxInt)
	return l
}

//Randomize re-inserts new random values into list
func (l list) Randomize(maxInt int) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < len(l); i++ {
		l[i] = random.Intn(maxInt)
	}
}

//Reverse the ordering of elements.
func (l list) Reverse() {
	tmp := make(list, len(l))
	for i := 0; i < len(l); i++ {
		tmp[i] = l[len(l)-i-1]
	}
	for i := 0; i < len(l); i++ {
		l[i] = tmp[i]
	}
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
