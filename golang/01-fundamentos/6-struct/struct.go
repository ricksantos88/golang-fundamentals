package main

import "fmt"

type user struct {
	name string
	age uint8
	address endereco
}

type endereco struct {
	street string
	number uint8
}


func main() {
	fmt.Println("Structs")

	var user1 user
	fmt.Println(user1)

	user1.name = "João"
	user1.age = 30
	fmt.Println(user1)

	address1 := endereco{"Rua dos Bobos", 0}
	user2 := user{"Maria", 25, address1}
	fmt.Println(user2)


}