package utils

import "regexp"

func RemoveSpecialCharacters(str string) string {
	// Compila a expressão regular que corresponde a qualquer caractere que não seja uma letra maiúscula ou minúscula, um número ou um espaço.
	// A função MustCompile irá panic se a expressão regular for inválida, mas isso é aceitável para uma função de utilidade.
	re := regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	// Substitui todos os caracteres que não se encaixam no padrão por uma string vazia.
	return re.ReplaceAllString(str, "")
}
