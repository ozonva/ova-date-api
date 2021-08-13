package utils

import "fmt"

func SplitSlice(input []int, size int) [][]int {
	var r [][]int

	var chunk = make([]int, 0)
	for i := range input {
		chunk = append(chunk, input[i])
		if ((i+1)%size == 0) || (len(input)-i-1 == 0) {
			r = append(r, chunk)
			chunk = make([]int, 0)
		}
	}

	return r
}

func ExcludeItems(input []string, filter map[string]bool) []string {
	var r []string

	for _, val := range input {
		if _, ok := filter[val]; !ok {
			r = append(r, val)
		}
	}

	return r
}

func ReverseKeysValues(input map[string]int) map[int]string {
	r := make(map[int]string)

	for ind, val := range input {
		r[val] = ind
	}

	return r
}

func PrintSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
