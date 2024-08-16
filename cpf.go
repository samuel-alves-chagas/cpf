package main

import (
	"strconv"
	"strings"
)

// ValidaCPF retorna se o cpf é válido ou não
func validaCPF(cpf string) bool {

	// TO DO: Validar que não tem nenhum traço nem ponto no cpf digitado e removê-los

	// To Do: Validar que o CPF só possui números

	// Validar que o cpf possui 11 caracteres
	if len(cpf) != 11 {
		return false
	}

	numerosCPF, _ := separaCPF(cpf)

	// Validar se os dígitos verificadores estão corretos
	digito1, digito2 := validaDigitosVerificadores(numerosCPF)

	return digito1 || digito2
}

func separaCPF(cpf string) ([]int, []string) {

	digitosCPF := strings.Split(cpf, "")
	numerosCPF := make([]int, 11)

	for indice, digitoAtual := range digitosCPF {
		numeroAtual, err := strconv.Atoi(digitoAtual)
		if err != nil {
			panic(err)
		}
		numerosCPF[indice] = numeroAtual
	}

	return numerosCPF, digitosCPF
}

func validaDigitosVerificadores(numerosCPF []int) (bool, bool) {
	multiplosDigito1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	multiplosDigito2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	somaParaPrimeiroDigito := multiplicaDigitosEMultiplos(numerosCPF[:9], multiplosDigito1)
	somaParaSegundoDigito := multiplicaDigitosEMultiplos(numerosCPF[:10], multiplosDigito2)

	digito1Calculado := somaParaPrimeiroDigito % 11
	digito2Calculado := somaParaSegundoDigito % 11

	if digito1Calculado == 10 {
		digito1Calculado = 0
	}

	if digito2Calculado == 10 {
		digito2Calculado = 0
	}

	digito1valido := digito1Calculado == numerosCPF[9]
	digito2valido := digito2Calculado == numerosCPF[10]

	return digito1valido, digito2valido
}

func multiplicaDigitosEMultiplos(digitos []int, multiplos []int) int {
	acumulador := 0
	for indice, valor := range digitos {
		acumulador = acumulador + valor*multiplos[indice]
	}

	return acumulador
}
