package main

import "fmt"

func main() {
	// TO DO: Pegar o cpf que o usuário digitou no terminal
	const cpf = ""

	// Validar se os dígitos verificadores estão corretos
	ehValido := validaCPF(cpf)
	fmt.Println(ehValido)
}
