package main

import "fmt"

func main() {
	get_fourth_element_of_array()
	slice_length()
	slice_by_array()
	min_element_in_array()
}

// #1.
func get_fourth_element_of_array() {
	arr := [5]int64{1, 2, 3, 4}
	fmt.Println(arr[3])
}

// #2.
func slice_length() {
	slice := make([]int, 3, 9)
	fmt.Println(len(slice) == 3)
}

// #3.
func slice_by_array() {
	x := [6]string{"a", "b", "c", "d", "e", "f"}
	res := [3]string{"c", "d", "e"}

	fmt.Println(x[2:5], res, len(x[2:5]) == len(res))
}

// #4.
func min_element_in_array() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	tmp := x[0]
	for i := 1; i < len(x); {
		if tmp < x[i] {
			i++
		} else {
			tmp = x[i]
			i++
		}
	}

	fmt.Println(tmp)
}
