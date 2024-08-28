package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var cpf string

	fmt.Println("Informe o seu CPF:")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		cpf = scanner.Text()
		break
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

	if validaCPF(cpf) {
		fmt.Println("O CPF informado é válido")
	} else {
		fmt.Println("O CPF informado não é válido")
	}
}
