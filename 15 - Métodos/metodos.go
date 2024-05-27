package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do Usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	fmt.Println("Métodos")
	usuario1 := usuario{"Raphael", 17}
	usuario1.salvar()
	maiorDeIdade1 := usuario1.maiorDeIdade()
	fmt.Println(maiorDeIdade1)

	usuario2 := usuario{"Thaina", 19}
	usuario2.salvar()
	maiorDeIdade2 := usuario2.maiorDeIdade()
	fmt.Println(maiorDeIdade2)
	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)

}
