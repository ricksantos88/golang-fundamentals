package main

import (
    "fmt"
	"math-module/mathutils"
)

func main() {
    soma := mathutils.Add(10, 5)
    subtracao := mathutils.Subtract(10, 5)
    produto := mathutils.Multiply(3.5, 2.0)

    // Variáveis para exibição
    resultado := fmt.Sprintf(
        "Soma: %d\nSubtração: %d\nProduto: %.2f",
        soma, subtracao, produto,
    )

    fmt.Println(resultado)
}