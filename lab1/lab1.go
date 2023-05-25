package main

import (
	initarr "algorithms/packages/arrays/init_arr"
	"algorithms/packages/arrays/search"
	"algorithms/packages/measure"
	"fmt"
	"log"
)

func insertAt(a []int, index int, value int) []int {
	if len(a) == index {
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...)
	a[index] = value
	return a
}

func removeAt(a []int, index int) []int {
	return append(a[:index], a[index+1:]...)
}

func main() {
	var err error

	var size int
	_, err = fmt.Scan(&size)
	if err != nil {
		log.Fatal(err)
	}

	ord_arr := initarr.InitOrd(size)
	fmt.Println(ord_arr)

	var target int
	_, err = fmt.Scan(&target)
	if err != nil {
		log.Fatal(err)
	}

	measure.Measure(func() {
		fmt.Printf("Binary search: %d\n", search.BinarySearch(ord_arr, target))
	})

	measure.Measure(func() {
		fmt.Printf("Simple search: %d\n", search.Search(ord_arr, target))
	})

	new_ord := initarr.InitOrd(10)
	new_rand := initarr.InitRand(10)
	fmt.Println(new_rand)
	fmt.Println(new_ord)

	measure.Measure(func() {
		new_ord = insertAt(new_ord, 5, 111)
		fmt.Print("ord ")
		fmt.Print(new_ord)
	})

	measure.Measure(func() {
		new_rand = insertAt(new_rand, 5, 111)
		fmt.Print("rand ")
		fmt.Print(new_rand)
	})

	measure.Measure(func() {
		new_ord = removeAt(new_ord, 5)
		fmt.Print("ord ")
		fmt.Print(new_ord)
	})

	measure.Measure(func() {
		new_rand = removeAt(new_rand, 5)
		fmt.Print("rand ")
		fmt.Print(new_rand)
	})
}
