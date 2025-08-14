package main

import "fmt"

type person struct {
	name    string
	surname string
	age     int
}

type student struct {
	person
	course     string
	university string
}

func main() {
	fmt.Println("Herança, só que não!")

	// Instanciando uma pessoa
	p1 := person{name: "John", surname: "Nolan", age: 40}
	fmt.Println("Person:", p1)
	
	// Instanciando um estudante
	s1 := student{p1, "ADS", "UNAMA"}
	fmt.Println(s1)

	// Apresentando o aluno e acessando propriedades da composição
	fmt.Printf("Aluno: %s %s, %d anos, curso %s na %s\n", s1.name, s1.surname, s1.age, s1.course, s1.university)

}
