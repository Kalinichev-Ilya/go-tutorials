package main

import (
	"fmt"
)

func main() {
	create()
	create_another_key()
	delete_element()
	big_map()
}

func create() {
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])
}

func create_another_key() {
	x := make(map[int]int)
	x[1] = 10
	fmt.Println(x[1])
}

func delete_element() {
	x := make(map[int]int)
	x[1] = 10
	fmt.Println(x[1])
	delete(x, 1)
	fmt.Println(x[1])

	name, ok := x[1]
	fmt.Println(name, ok)

	// Сперва мы пробуем получить значение из карты, а затем,
	// если это удалось, мы выполняем код внутри блока.
	if name, ok := x[1]; ok {
		fmt.Println(name, ok)
	}

}

func big_map() {

	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "solid",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": map[string]string{
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": map[string]string{
			"name":  "Neon",
			"state": "gas",
		},
	}

	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}


