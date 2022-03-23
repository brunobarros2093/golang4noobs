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
	//paralelismo nos testes
	t.Parallel()
	// vários cenários
	cenariosDeTestes := []cenarioDeTest{

		{"Rua John Doe", "Rua"},
		{"Avenida John Doe", "Avenida"},
		{"Praça John Doe", "Tipo Inválido"},
		{"Estrada John Doe", "Estrada"},
		{"Rodovia John Doe", "Rodovia"},
		{"", "Tipo Inválido"},
		{"birl", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTestes {
		retornoRecebido := TipoDeEndereco(cenario.endereco)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("Tipo recebido é %s é diferente do esperado %s", retornoRecebido, cenario.retornoEsperado)
		}
	}
}

//TestQualquer para testar paralelismo
func TestQualquer(t *testing.T) {
	//paralelismo nos testes
	t.Parallel()
	if 1 > 2 {
		t.Errorf("Teste Quebrou oh no ")
	}
}
