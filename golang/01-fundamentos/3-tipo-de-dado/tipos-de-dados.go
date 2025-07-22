package main

import (
	"errors"
	"fmt"
)

func main() {
	var numeroInt8 int8 = 127 // -128 to 127
	fmt.Println(numeroInt8)

	var numeroInt16 int16 = 32767 // -32768 to 32767
	fmt.Println(numeroInt16)

	var numeroUint8 uint8 = 255 // 0 to 255
	fmt.Println(numeroUint8)

	//Alias
	var numeroRune rune = 123456 // int3
	fmt.Println(numeroRune)

	// Números Reais
	var numeroFloat32 float32 = 123.45
	fmt.Println(numeroFloat32)

	var numeroFloat64 float64 = 1.7976931348623157e+308 // Maximum float64 value
	fmt.Println(numeroFloat64)

	// String
	var texto string = "Texto"
	fmt.Println(texto)

	texto2 := "Texto 2"
	fmt.Println(texto2)

	char := 'A'
	fmt.Println(char)

	var booleano bool = true
	fmt.Println(booleano)

	var err error = errors.New("Erro")
	fmt.Println(err)
}
