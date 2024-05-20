package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Excrever()
	erro := checkmail.ValidateFormat("raphael.segalla#email.com")
	fmt.Println(erro)
}
