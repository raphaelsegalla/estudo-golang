package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de partida")

	// #Subistuido pelo codigo abaixo#
	// aplicacao := app.Gerar()
	// erro := aplicacao.Run(os.Args)
	// if erro != nil {
	// 	log.Fatal(erro)
	// }

	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}

}
