package main

import "fmt"

func main() {
	variavel1 := 10
	variavel2 := variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++

	fmt.Println(variavel1, variavel2)

	// ponteiros -> referência de memória

	var	variavel3 int
	var ponteiro *int

	fmt.Println(variavel3, ponteiro)

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3, ponteiro)
	fmt.Println(variavel3, *ponteiro) // o  * é usado para "desreferenciar" o ponteiro

	variavel3 = 150
	fmt.Println(variavel3, *ponteiro)
}