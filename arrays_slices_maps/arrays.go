package main

import "fmt"

func main() {
	structure()
	average_value()
	average_value_ref()
	average_value_for_loop()
	average_value_array()
}

func structure() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)
}

func average_value() {
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64 = 0
	for i := 0; i < 5; i++ {
		total += x[i]
	}
	fmt.Println(total / 5)
}

func average_value_ref() {
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64 = 0
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	fmt.Println(total / float64(len(x)))
}

func average_value_for_loop() {
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64 = 0
	for _, val := range x {
		total += val
	}
	fmt.Println(total / float64(len(x)))
}

func average_value_array() {
	x := [5]float64{98, 93, 77, 82, 83}

	var total float64 = 0
	for _, val := range x {
		total += val
	}
	fmt.Println(total / float64(len(x)))
}
