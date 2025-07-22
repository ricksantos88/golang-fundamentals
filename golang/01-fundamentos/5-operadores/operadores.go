package main

import "fmt"

func main() {

	// Operadores aritméticos
	sum := 2 + 2
	subt := 2 - 2
	div := 2 / 2
	mult := 2 * 2

	fmt.Println(sum, subt, div, mult)

	// Operadores de atribuição
	var val1 string = "Hello"
	val2 := "World"
	fmt.Println(val1, val2)

	// operadores de relacionais
	fmt.Println(2 > 1)
	fmt.Println(2 < 1)
	fmt.Println(2 == 2)
	fmt.Println(2 != 2)
	fmt.Println(2 >= 2)
	fmt.Println(2 <= 2)

	// Operadores lógicos
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	//operadores unários
	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)

	numero--
	numero -= 20
	fmt.Println(numero)

	numero *= 3
	fmt.Println(numero)

	numero /= 2
	fmt.Println(numero)
	
	numero %= 2
	fmt.Println(numero)

}