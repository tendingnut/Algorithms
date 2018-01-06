//Implementation of common sorting algorithms for practice, without
//using package sort from the standard library
package sorts

import (
	"math"
	"math/rand"
	"time"
)

type list []int

//Newlist is a constructor for list.
func NewList(size, maxInt int) list {
	l := make(list, size)
	l.Randomize(size, maxInt)
	return l
}

//Randomize re-inserts new random values into list
func (l list) Randomize(size, maxInt int) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < size; i++ {
		l[i] = random.Intn(maxInt)
	}
}

//Reverse the ordering of elements.
func (l list) Reverse() {
	tmp := list{}
	for i := 0; i < len(l); i++ {
		tmp[i] = l[len(l)-i-1]
	}
}

//Insertion sort has n^2 complexity.
func Insertion(slice []int) []int {
	list := make([]int, len(slice))
	copy(list, slice)
	for j := 1; j < len(list); j++ {
		key := list[j]
		i := j - 1
		for i >= 0 && list[i] > key {
			list[i+1] = list[i]
			i--
		}
		list[i+1] = key
	}
	return list
}

//Merge sort has nlg(n) complexity.
func Merge(slice []int) []int {
	list := make([]int, len(slice))
	copy(list, slice)
	mergeSort(list, 0, len(list)-1)
	return list
}

//mergeSort divides the list recursively into sub-sections by calling mergeSort(),
//and calls 'merge()' on each sub-section (p..r) to sort.
func mergeSort(list []int, p, r int) { //p is first element, r is last element
	if p < r { // if more than 1 element in list, otherwise already sorted
		q := int(math.Floor(float64((p + r) / 2))) // middle element used to split list
		mergeSort(list, p, q)
		mergeSort(list, q+1, r)
		merge(list, p, q, r)
	}
}

//called on a sub-section of list, where each half is sorted (p..q) and (q+1..r),
//and put into separate slices, which are used to compare and sort entire sub-section (p..r)
func merge(list []int, p, q, r int) {
	var A, B []int
	//loops fill up slices with appropriate values from list
	for i := p; i <= q; i++ {
		A = append(A, list[i])
	}
	for i := q + 1; i <= r; i++ {
		B = append(B, list[i])
	}
	A = append(A, math.MaxInt64) //add sentinel value to end of each slice
	B = append(B, math.MaxInt64)
	i, j := 0, 0
	//A and B are sorted, so loop compares and adds lowest of each slice
	//to p..r in list
	for k := 0; k <= r-p; k++ {
		if A[i] <= B[j] {
			list[p+k] = A[i]
			i++
		} else {
			list[p+k] = B[j]
			j++
		}
	}
}
