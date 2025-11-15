//calculadora simples
//calcular operacões simples (+, -, *, /)
// perguntar ao usuario qual operadoração usar
//exibir o resultado

package main

import "fmt"

func opCalcular(op string, a, b float64) float64 { //para calcular a operação no que o usuario escolher
	if op == "+" {
		return a + b
	}
	if op == "-" {
		return a - b
	}
	if op == "*" {
		return a * b
	}
	if op == "/" {
		return a / b
	}
	return 0 // caso não exista um desses operadores
}

func layout () {
	for { //loop infinito (tem que parar com um comando)
		var operador string
		var num1, num2 float64
	
		fmt.Println ("Escolha uma operação: + - * /")

		fmt.Scan(&operador)

		//se o usuario quiser sair, digita a palavra "sair"
		if operador == "sair" {
			fmt.Println("Programa encerrado!")
			break // e para (sai do loop)
		}
		//criar uma mensagem para quando o usuario digiar um operador invalido
		if operador != "+" && operador != "-" && operador != "*" && operador != "/" {
			fmt.Println("Operador invalido! use: + - * /, para continuar")
			continue //volta para o loop
		}

		fmt.Println("Digite o primeiro número:")
		fmt.Scan(&num1)

		fmt. Println("Digite o Segundo número:")
		fmt.Scan(&num2)

		resultado := opCalcular(operador, num1, num2)

		fmt.Printf("Resultado: %.2f %s %.2f = %.2f\n", num1, operador, num2, resultado)
}
}

func main() {

	layout()
	
}