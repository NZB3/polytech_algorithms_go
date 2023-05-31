package main

import (
	initarr "alg/pkgs/arrays/init_arr"
	"alg/pkgs/arrays/search"
	"alg/pkgs/arrays/useless"
	"alg/pkgs/measure"
	"fmt"
	"log"
)

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
		new_ord = useless.InsertAt(new_ord, 5, 111)
		fmt.Print("ord ")
		fmt.Print(new_ord)
	})

	measure.Measure(func() {
		new_rand = useless.InsertAt(new_rand, 5, 111)
		fmt.Print("rand ")
		fmt.Print(new_rand)
	})

	measure.Measure(func() {
		new_ord = useless.RemoveAt(new_ord, 5)
		fmt.Print("ord ")
		fmt.Print(new_ord)
	})

	measure.Measure(func() {
		new_rand = useless.RemoveAt(new_rand, 5)
		fmt.Print("rand ")
		fmt.Print(new_rand)
	})

	measure.Measure(func() {
		new_ord = useless.Append(new_ord, 111)
		fmt.Print("ord ")
		fmt.Print(new_ord)
	})

	measure.Measure(func() {
		new_rand = useless.Append(new_rand, 111)
		fmt.Print("rand ")
		fmt.Print(new_rand)
	})

	measure.Measure(func() {
		new_ord = useless.Prepend(new_ord, 111)
		fmt.Print("ord ")
		fmt.Print(new_ord)
	})

	measure.Measure(func() {
		new_rand = useless.Prepend(new_rand, 111)
		fmt.Print("rand ")
		fmt.Print(new_rand)
	})
}
