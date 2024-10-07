package main

import "fmt"

func merge([]int a, []int b) []int {
	i :=0
	j :=0
	
	res := []int{} 
	for i < len(a) && j < len(b) {
        if a[i] <b[j] {
			res = append(res,a[i])
			i++
		}

		if b[j] < a[i] {
			res = append(res, b[j])
			j++
		}
	}
	for i < len(a) {
		res = append(res, a[i])
		i++
	}

	for j < len(b) {
		res = append(res, b[j])
		j++
	}
	return res
}

func mergesort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	first := mergesort(arr[:len(arr)/2])
    second := mergesort(arr[len(arr)/2:])
	return merge(first, second)
}
func main() {
	unsorted := []int{3,1,2,5,4}
	sorted := mergesort(unsorted)
}