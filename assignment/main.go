package main

import "fmt"

func main() {
	var testvalue []int

	for i := 0; i < 11; i++ {
		testvalue = append(testvalue, i)
	}

	for _, value := range testvalue {
		if (value % 2) == 0 {
			fmt.Println(value, " is even")
		} else {
			fmt.Println(value, " is odd")
		}
	}
}
