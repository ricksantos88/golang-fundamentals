package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func mathCalcs(a int, b int) (int, int, int, int) {
	sum := a + b
	sub := a - b
	mult := a * b
	div := a / b

	return sum, sub, mult, div
}

func main() {
	soma := sum(1, 2)
	fmt.Println(soma)

	var f = func() {
		fmt.Println("Função sem nome")
	}
	f()

	var f2 = func(param string) {
		fmt.Println(param)
	}
	f2("Função com parâmetro")

	var f3 = func(param string) string {
		fmt.Println(param)
		return "retorno da função"
	}

	retornoDaFuncao := f3("Função com retorno")
	println(retornoDaFuncao)

	sum, sub, mult, div := mathCalcs(10, 2)
	fmt.Println(sum, sub, mult, div)

	fmt.Println(mathCalcs(10, 2))
}
