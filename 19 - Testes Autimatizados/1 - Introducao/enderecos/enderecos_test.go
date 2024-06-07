package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {

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

	// if tipoDeEndercoRecebido != tipoDeEnderecoEsperado {
	// 	// t.Error("O tipo recebido é diferente do esperado!")
	// 	t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s",
	// 		tipoDeEnderecoEsperado,
	// 		tipoDeEndercoRecebido,
	// 	)
	// }
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