package main

import (
	initarr "alg/pkgs/arrays/init_arr"
	"alg/pkgs/arrays/search"
	"alg/pkgs/arrays/useless"
	"alg/pkgs/linkedlists/doublelinkedlist"
	"alg/pkgs/linkedlists/linkedlist"
	"alg/pkgs/measure"
	"fmt"
	"log"
)

func main() {
	fmt.Print("Enter size of an array: ")

	var size int
	_, err := fmt.Scan(&size)
	if err != nil {
		log.Fatal(err)
	}
	var arr []int
	list := linkedlist.NewLinkedList[int]()
	dlist := doublelinkedlist.NewDoubleLinkedList[int]()

	measure.Measure(func() {
		arr = initarr.InitRand(size)
		fmt.Println(arr)
	})

	measure.Measure(func() {
		list.Fill(arr)
		list.Print()
	})

	measure.Measure(func() {
		dlist.Fill(arr)
		dlist.Print()
	})

	fmt.Println("append ")
	measure.Measure(func() {
		arr = useless.Append(arr, 111)
		fmt.Print(arr)
	})

	measure.Measure(func() {
		list.Append(111)
		list.Print()
	})

	measure.Measure(func() {
		dlist.Append(111)
		dlist.Print()
	})

	fmt.Println("Prepend")
	measure.Measure(func() {
		arr = useless.Prepend(arr, 111)
		fmt.Print(arr)
	})

	measure.Measure(func() {
		list.Prepend(111)
		list.Print()
	})

	measure.Measure(func() {
		dlist.Prepend(111)
		dlist.Print()
	})

	fmt.Print("Enter value to insert: ")
	var value int
	_, err = fmt.Scan(&value)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Enter index where to insert: ")
	var index int
	_, err = fmt.Scan(&index)
	if err != nil {
		log.Fatal(err)
	}

	measure.Measure(func() {
		arr = useless.InsertAt(arr, index, value)
		fmt.Print("arr ")
		fmt.Print(arr)
	})

	measure.Measure(func() {
		list.Set(index, value)
		list.Print()
	})

	measure.Measure(func() {
		dlist.Set(index, value)
		dlist.Print()
	})

	fmt.Println("Remove")

	measure.Measure(func() {
		arr = useless.RemoveAt(arr, index)
		fmt.Print(arr)
	})

	measure.Measure(func() {
		list.Remove(index)
		list.Print()
	})

	measure.Measure(func() {
		dlist.Remove(index)
		dlist.Print()
	})

	fmt.Println("Enter size of an array: ")
	_, err = fmt.Scan(&size)
	if err != nil {
		log.Fatal(err)
	}

	ord_arr := initarr.InitOrd(size)
	new_list := linkedlist.NewLinkedList[int]()
	new_dlist := doublelinkedlist.NewDoubleLinkedList[int]()
	go func() {
		fmt.Println(ord_arr)
	}()
	go func() {
		new_list.Fill(ord_arr)
	}()
	go func() {
		new_dlist.Fill(ord_arr)
	}()

	fmt.Print("Enter value to find: ")
	var target int
	_, err = fmt.Scan(&target)
	if err != nil {
		log.Fatal(err)
	}

	measure.Measure(func() {
		fmt.Printf("Binary search: %d\n", search.BinarySearch(ord_arr, target))
	})

	measure.Measure(func() {
		fmt.Println(new_list.Find(target))
	})

	measure.Measure(func() {
		fmt.Println(new_dlist.Find(target))
	})

}
