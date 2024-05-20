package main

import "fmt"

func main() {
	fmt.Println("Funções anônimas")

	retorno := func(texto string) string {
		fmt.Println(texto)
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando Parâmetro")

	fmt.Println(retorno)
}
