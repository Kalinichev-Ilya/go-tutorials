package main

import (
	"fmt"
)

func main() {
	create_slice()
	slice_funk_append()
	slice_funk_copy()
}

func create_slice() {
	var x []float64
	y := make([]float64, 5)
	z := make([]float64, 5, 10)
	fmt.Println(x, y, z)

	//low and high
	arr := [5]float64{1, 2, 3, 4, 5}
	a := arr[0:5]
	b := arr[1:4]
	c := arr[1:len(arr)]
	fmt.Println(arr, a, b, c)

	d := arr[0:]
	e := arr[:5]
	f := arr[0:5]
	g := arr[:]
	k := arr[0:len(arr)]
	fmt.Println(d, e, f, g, k)
}

func slice_funk_append() {
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)
}

func slice_funk_copy() {
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	fmt.Println(slice1)  // 1 2 3
	fmt.Println(slice2)  // 0 0
	copy(slice2, slice1) // slice2 = 1 2
	fmt.Println(slice1, slice2)
}
