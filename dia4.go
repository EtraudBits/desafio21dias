package main

//codigo referente a escolha do melhor Ar condicionado para o ambiente

import "fmt"

//função fora do escopo da main para a area do ambiente do ar
func arIdeal(area float64) string {

	
	//usando condicionais if/else

	if area >=1 && area <=12 {
		return "7000 btus"
	}
	if area >=13 && area <=17 {
		return "9000 btus"
	}
	if area >=18 && area <=25 {
		return "12000 btus"
	}
	if area > 25 {
		return "Usa um aparelho maior que 12000 btus"
	}
	//senão (caso não esteja em nunhuma das alternativa) retorna mensagem de erro
	return "área inválida, por favor informe uma área válida."
}

func main () {

	//declaração de variaveis para calcular o metro quadrado
	var comprimento float64
	var largura float64

	//pede para o usuario digitar o comprimento e a largura do vão
	fmt.Println ("Qual o Comprimento: ")
	fmt.Scan(&comprimento)

	fmt.Println ("Qual a Largura: ")
	fmt.Scan(&largura)

	//calcular o metro quadrado

	vaoTotal := comprimento*largura

	resultado := arIdeal(vaoTotal)

	//caso o resultado não for valido
	if resultado == "área inválida, por favor informe uma área válida." {
		fmt.Println(resultado)
		return //e para
	}
	// exibe a resposta de qual ar será o ideal para o ambiente digitado
	//%.2f força exibir apenas 2 digitos após o ponto
	fmt.Printf("O ar-condicionado ideal para seu ambiente de %.2f m² é: %s", vaoTotal, resultado)
}