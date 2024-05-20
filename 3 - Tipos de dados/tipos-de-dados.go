package main

import (
	"errors"
	"fmt"
)

func main() {
	// NÙMEROS INTEIROS

	var numero int64 = -100000000000
	fmt.Println(numero)

	var numero2 uint32 = 1000
	fmt.Println(numero2)

	// alias
	// INT32 = RUNE
	var numero3 rune = 123
	fmt.Println(numero3)

	// BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	// FIM NÙMEROS INTEIROS

	// NÙMEROS REAIS

	var numeroReal1 float32 = 1230000000.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 12345.67
	fmt.Println(numeroReal2)

	// FIM NÙMEROS REAIS

	// STRINGS

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Text2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)

	// FIM STRINGS

	var boleano1 bool = true
	fmt.Println(boleano1)

	var erro error
	fmt.Println(erro)

	var erro1 error = errors.New("Erro interno")
	fmt.Println(erro1)
}
