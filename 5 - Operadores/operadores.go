package main

import "fmt"

func main() {
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 2
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 25

	soma2 := numero1 + numero2

	fmt.Println(soma2)

	// ATRIBUICAO
	var variavel1 string = "String"
	variavel2 := "String"

	fmt.Println(variavel1, variavel2)
	// FIM ATRIBUICAO

	// RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)
	// FIM RELACIONAIS

	// LOGICOS
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	// FIM LOGICOS

	// UNARIOS
	numero := 10
	numero++
	numero += 15 // numero = numero + 15

	fmt.Println(numero)

	numero--

	fmt.Println(numero)

	// FIM UNARIOS
}
