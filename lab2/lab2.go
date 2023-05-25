package main

import (
	initarr "algorithms/packages/arrays/init_arr"
	"algorithms/packages/arrays/sorts"
	"algorithms/packages/measure"
	"fmt"
	"log"
)

func main() {
	var err error
	var n int
	_, err = fmt.Scan(&n)
	if err != nil {
		log.Fatal(err)
	}

	arr := initarr.InitRand(n)
	fmt.Println(arr)

	//measure.Measure(func() {
	//sorts.BubleSort(arr)
	//})

	measure.Measure(func() {
		sorts.InsertSort(arr)
	})

	measure.Measure(func() {
		sorts.SelectSort(arr)
	})

	measure.Measure(func() {
		sorts.QuickSort(arr)
	})

	measure.Measure(func() {
		sorts.MergeSort(arr)
	})
}
