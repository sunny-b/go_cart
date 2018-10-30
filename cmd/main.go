package main

import "fmt"

func main() {
	slice := make([]int, 0)
	slice = append(slice, 1)
	slice = append(slice, 2)
	fmt.Println(len(slice), slice[0], slice[1])
}
