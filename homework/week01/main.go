package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Before deletion, sliec = ", slice)
	slice = Remove(slice, 5)
	fmt.Println("After deletion, sliec = ", slice)
}
