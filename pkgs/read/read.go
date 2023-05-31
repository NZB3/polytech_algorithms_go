package read

import (
	"fmt"
	"log"
)

func ReadIntArr() []int {
	var count int
	_, err := fmt.Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	var arr = make([]int, count)
	for i := 0; i < len(arr); i++ {
		_, err := fmt.Scan(&arr[i])
		if err != nil {
			log.Fatal(err)
		}
	}
	return arr
}

func ReadStrArr() []string {
	var count int
	_, err := fmt.Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	var arr = make([]string, count)
	for i := 0; i < len(arr); i++ {
		_, err := fmt.Scan(&arr[i])
		if err != nil {
			log.Fatal(err)
		}
	}
	return arr
}
