package enderecos

import (
	"testing"
)

type cenarioDeTest struct {
	endereco        string
	retornoEsperado string
}

// TestTipoDeEndereco testes unitários - recebe um ponteiro de T
func TestTipoDeEndereco(t *testing.T) {
	// vários cenários
	cenariosDeTestes := []cenarioDeTest{

		{"Rua John Doe", "Rua"},
		{"Avenida John Doe", "Avenida"},
		{"Praça John Doe", "Tipo Inválido"},
		{"Estrada John Doe", "Estrada"},
		{"Rodovia John Doe", "Rodovia"},
		{"", "Tipo Inválido"},
		{"Rua", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTestes {
		retornoRecebido := TipoDeEndereco(cenario.endereco)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("Tipo recebido é %s é diferente do esperado %s", retornoRecebido, cenario.retornoEsperado)
		}
	}
}
