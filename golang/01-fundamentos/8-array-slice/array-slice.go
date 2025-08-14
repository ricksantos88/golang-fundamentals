package main

import "fmt"

func main() {
	fmt.Println("Array and Slices")

	var array1 [5]int
	fmt.Println(array1)

	array1[0] = 10
	fmt.Println(array1)

	array2 := [5]string{"p1", "p2", "p3", "p4", "p5"}
	fmt.Println(array2)


	array3 := [...]int{1,2,3,4,5}
	fmt.Println(array3)
}