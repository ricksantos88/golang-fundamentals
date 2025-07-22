package main

import (
	"fmt"
	"module/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Writing from file main.go ...")
	auxiliar.Write()
	error := checkmail.ValidateFormat("henrique.santos1988@gmail.com")
	fmt.Println(error)
}
