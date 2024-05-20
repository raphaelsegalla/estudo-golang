package main

import "fmt"

type usuario struct {
	nome     string
	idade    int8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	fmt.Println(u)

	u.nome = "Davi"
	u.idade = 20

	fmt.Println(u)

	usuario2 := usuario{"Davi", 21, endereco{"rua", 2}}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Davi"}
	fmt.Println(usuario3)
}
