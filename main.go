package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Print("\nSeja bem vindo ao validador do CPF, informe a função desejada:\n\n1 - Validar CPF\n9 - Sair\n")

	var text string
	reader := bufio.NewReader(os.Stdin)

LeituraDados:
	for {
		text, _ = reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		switch text {
		case "1":
			fmt.Print("\nInforme o CPF que deve ser validado: ")
			text, _ = reader.ReadString('\n')
			text = strings.Replace(text, "\n", "", -1)

			if validaCPF(text) {
				fmt.Print("O CPF informado é válido\n")
			} else {
				fmt.Print("O CPF informado NÃO É VÁLIDO!\n")
			}
		case "9":
			break LeituraDados
		default:
			fmt.Print("As únicas funções possíveis são:\n\n1 - Validar CPF\n9 - Sair\n")
		}

		fmt.Print("\nPara validar outro CPF pressione 1, para sair pressione 9\n")
	}
}
