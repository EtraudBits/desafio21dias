package main

//elaborar um cronômetro simples

import (
"fmt"
"time"
)

func cronometro (segundos int) {
	for i:= 1; i <= segundos; i++ {
		fmt.Println("Tempo:", i, "Segundo(s)")
		time.Sleep(1 * time.Second) //para aguardar 1 segundo e continua
	}

	fmt.Println("Acabou o Tempo!")
}

func main() {

	var tempo int

	fmt.Println("Qual o tempo para o cronometro?")
	fmt.Scan(&tempo) //&chama a variavel

	if tempo <= 0 {
		fmt.Println("Tempo permitido maior que 0!")
		return
	}

	cronometro(tempo)
}