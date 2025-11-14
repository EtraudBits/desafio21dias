package main

import "fmt" // pacote para imprimir e ler dados

//missão dia 3 - estudar variaveis, tipo de dados e operadores
//objetivo criar um script simples que receba entrada e exiba resultado

func main ()  {

	//variaveis simples
	var nome string
	var idade int

	fmt.Println("Digite seu nome:") //exibe na tela para o usuario
	fmt.Scanln(&nome) // >Scanln< serve para ler o que o usuario digitou

	fmt.Println("Digite sua idade:")
	fmt.Scanln(&idade)

	//usando agora operadores

	tempoDev := idade + 1

	//exibi o resultado
	fmt.Println("Olá,", nome, ".")
	fmt.Println("Você será um DEV com,", tempoDev, "anos")

	fmt.Println("Continue com foco, ",nome, ", você vai conseguir!")
	
}