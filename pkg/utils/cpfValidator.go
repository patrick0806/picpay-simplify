package utils

import "strings"

func IsValidCPF(cpf string) bool {
	cpf = strings.Replace(cpf, ".", "", -1)
	cpf = strings.Replace(cpf, "-", "", -1)

	if len(cpf) != 11 {
		return false
	}

	// Primeiro dígito
	sum := 0
	for i := 0; i < 9; i++ {
		sum += int(cpf[i]-'0') * (10 - i)
	}
	d1 := 11 - sum%11
	if d1 >= 10 {
		d1 = 0
	}

	// Segundo dígito
	sum = 0
	for i := 0; i < 10; i++ {
		sum += int(cpf[i]-'0') * (11 - i)
	}
	d2 := 11 - sum%11
	if d2 >= 10 {
		d2 = 0
	}

	return int(cpf[9]-'0') == d1 && int(cpf[10]-'0') == d2
}
