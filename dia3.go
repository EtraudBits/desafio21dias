package main

import "fmt" //pacote para imprimir e ler dados

//Missão do dia 3 do desafio, estudar variaveis, tipos de dados e operadores
// objetivo criar um script simples que receba entradas e exiba resultados

func main() {

	//variaveis simples
	var nome string
	var idade int

	fmt.Println("Digite seu nome:") //Pergunta para o usuario 
	fmt.Scanln(&nome) //>Scanln< ler o que o usuario digitou

	fmt.Println("Digite sua Idade")
	fmt.Scanln(&idade)

	//usando operadores
	//vamos somar quantos anos ele terá em determinado tempo

	idadeFutura := idade + 1 //somará a idade + 1

	//vamos imprimir o resultado
	fmt.Println("Olá,", nome, ".")
	fmt.Println("Você será um DEV com", idadeFutura, "anos.")

	fmt.Println("Vamos, continue tento foco,", nome, ". Você conseguirá.")

}