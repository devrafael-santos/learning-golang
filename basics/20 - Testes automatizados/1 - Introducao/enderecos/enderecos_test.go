// TESTE DE UNIDADE

package enderecos_test

import (
	. "introducao-testes/enderecos"
	"testing"
)

type cenarioDeTest struct {
	enderecoInserido string
	enderecoEsperado string
}

func TestTipoDeEndereco(t *testing.T) {
	// Test(nome da funcao a ser testada)

	cenariosDeTeste := []cenarioDeTest{
		{"Rua ABS", "Rua"},
		{"Travessa ABC", "Travessa"},
		{"Avenida Paulista", "Avenida"},
		// { "Praça das Rosas", "Tipo inválido"},
		{"Estrada AAA", "Estrada"},
		{"RODOVIA ALEATORIA", "Rodovia"},
		// {"", "Tipo inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.enderecoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado: %s ", tipoDeEnderecoRecebido, cenario.enderecoEsperado)
		}
	}

}
