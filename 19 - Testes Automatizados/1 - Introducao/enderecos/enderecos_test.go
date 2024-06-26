package enderecos_test

import (
	. "introducao-testes/enderecos"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {

	t.Parallel()

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Praça das Rosas", "Tipo Inválido"},
		{"Estrada Qualquer", "Estrada"},
		{"RUA DOS BOBOS", "Rua"},
		{"AVENIDA REBOUÇAS", "Avenida"},
		{"", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				retornoRecebido,
				cenario.retornoEsperado,
			)
		}
	}
}

func TestQualquer(t *testing.T) {

	t.Parallel()

	if 1 > 2 {
		t.Error("Teste quebrou!")
	}
}

// func TestTipoDeEndereco(t *testing.T) {
// 	enderecoParaTeste := "Avenida Paulista"
// 	tipoDeEnderecoEsperado := "Avenida"
// 	tipoDeEndercoRecebido := TipoDeEndereco(enderecoParaTeste)

// 	if tipoDeEndercoRecebido != tipoDeEnderecoEsperado {
// 		// t.Error("O tipo recebido é diferente do esperado!")
// 		t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s",
// 			tipoDeEnderecoEsperado,
// 			tipoDeEndercoRecebido,
// 		)
// 	}
// }
