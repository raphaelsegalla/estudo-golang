package main

import "fmt"

func main() {
	canal := make(chan string, 2)
	canal <- "Olá Mundo!"
	canal <- "Programando em GO!"
	// canal <- "Terceira mensagem" // ocorre deadlock

	close(canal)

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	// mensagem := <-canal
	// mensagem1 := <-canal
	// fmt.Println(mensagem)
	// fmt.Println(mensagem1)
}
