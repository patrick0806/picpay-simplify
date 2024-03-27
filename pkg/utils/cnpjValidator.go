package utils

func IsValidCNPJ(cnpj string) bool {
	cnpj = RemoveSpecialCharacters(cnpj)

	if len(cnpj) != 14 {
		return false
	}

	//TODO - implement calculate check digit
	return true
}
