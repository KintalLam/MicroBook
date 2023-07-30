package main

import "fmt"

func Remove(slice []int, index int) []int {

	if index < 0 {
		fmt.Println("Index can not < 0")
		return slice
	} else if index >= len(slice) {
		fmt.Println("Index can not >= ", len(slice))
		return slice
	}

	//Header
	//header := slice[0:index]
	//fmt.Println("Header = ", header)
	//Trailer
	// trailer := slice[index+1:]
	// fmt.Println("Trailer = ", trailer)
	//Combination
	fmt.Println("Header = ", slice[0:index])
	fmt.Println("Trailer = ", slice[index+1:])
	// combination := append(slice[0:index], slice[index+1:]...)
	return append(slice[0:index], slice[index+1:]...)

}
