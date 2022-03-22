package enderecos

import "strings"

// TipoDeEndereco retorna se o tipo de endereço passado é válido e o retorna
func TipoDeEndereco(enderecoCompleto string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}
	endeToLower := strings.ToLower(enderecoCompleto)
	primeiraPalavra := strings.Split(endeToLower, " ")[0]
	//rua do krl
	// ["rua", "do", "krl"] -> pegamos a primeria palavra
	enderecoTemUmTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavra {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return strings.ToTitle(primeiraPalavra)
	}
	return "Tipo Inválido"

}
