package main

import (
	"fmt"
	"strings"
)

func main() {
	const value = 10
	_ = value
	fmt.Println("Done")

	var ch rune = '\a'
	fmt.Println(ch)

	// Array Demo

	var x1 [5]int
	var x2 [0]int
	var x3 [0]string

	var arr [3]int = [3]int{1, 2, 3}
	var arr2 = [3]int{1, 2, 3}
	arr3 := [3]int{1, 2, 3}

	//arr3 = [2]int{1,2}

	arr5 := [...]int32{1, 2, 3, 4, 5}

	fmt.Println(arr, arr2, arr3, arr5, x1, x2, x3, len(arr5))

	var list = []int64{1, 2, 3, 4, 5}
	l := len(list)
	c := cap(list)

	fmt.Println(list, l, c)

	list = make([]int64, 0, 5) //allocates an underlying array of size 10 and returns a slice of length 0
	l = len(list)
	c = cap(list)

	fmt.Println(list, l, c)

	list = make([]int64, 5, 5) //allocates an underlying array of size 10 and returns a slice of length 5
	l = len(list)
	c = cap(list)

	fmt.Println(list, l, c)

	list = append(list, 1)
	list = append(list, 2)
	list = append(list, 3)
	list = append(list, 4)
	list = append(list, 5)
	l = len(list)
	c = cap(list)

	fmt.Println(list, l, c)

	list = append(list, 6)

	l = len(list)
	c = cap(list)

	fmt.Println(list, l, c)

	// Map Demo

	var m1 map[int32]bool

	m1 = map[int32]bool{}

	m1[1] = false

	var m2 map[string]string

	m3 := make(map[int]int)

	ages := map[string]int32{
		"KKK": 32,
		"CCC": 23,
	}
	ages["RRR"] = 5

	ages["KKK"]++

	fmt.Println(m1, m2, m3, ages)

	str_map := strings.Map(func(r rune) rune {
		return r + 1
	}, "SDW")

	fmt.Println(str_map)
}
