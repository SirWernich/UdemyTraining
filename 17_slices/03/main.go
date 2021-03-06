package main

import "fmt"

func main() {
	mySlice := make([]int, 0, 16) // length 0, capacity 5

	fmt.Println("--------------")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println("--------------")

	for i := 0; i < 81; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len:", len(mySlice), "Capacity:", cap(mySlice),
			"Value:", mySlice[i], "first address:\t", &mySlice[0])
	}
}
