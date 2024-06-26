package enderecos

import "strings"

func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "travessa", "avenida", "estrada", "rodovia"}

	primeiraPalavraDoEndereco := strings.Split(endereco, " ")[0]
	primeiraPalavraDoEndereco = strings.ToLower(primeiraPalavraDoEndereco)

	for _, tipo := range tiposValidos {
		if primeiraPalavraDoEndereco == tipo {
			return strings.Title(tipo)
		}
	}
	return "Tipo inválido"
}
