package main

import (
	"fmt"
	"strings"
)

// 1) Sequência de Fibonacci
func pertenceFibonacci(num int) {
	a, b := 0, 1

	for b < num {
		a, b = b, a+b
	}

	if b == num || num == 0 {
		fmt.Println(num, "pertence à sequência de Fibonacci.")
	} else {
		fmt.Println(num, "não pertence à sequência de Fibonacci.")
	}
}

// 2) Verificação e contagem da letra 'a'
func analisaLetraA(s string) {
	count := 0
	for _, char := range s {
		if strings.ToLower(string(char)) == "a" {
			count++
		}
	}
	if count > 0 {
		fmt.Printf("A letra 'a' aparece %d vezes na string.\n", count)
	} else {
		fmt.Println("A letra 'a' não está presente na string.")
	}
}

// 3) Cálculo da variável SOMA
func calcularSoma() int {
	indice := 12
	soma := 0
	k := 1
	for k < indice {
		k = k + 1
		soma = soma + k
	}
	return soma
}

func imprimeSoma() {
	soma := calcularSoma()
	fmt.Printf("O valor da variável SOMA ao final do processamento é: %d\n", soma)
}

func main() {
	// 1) Sequência de Fibonacci
	num := 21 // Você pode alterar esse valor para testar diferentes números.
	pertenceFibonacci(num)
	pertenceFibonacci(num + 1)

	// 2) Verificação e contagem da letra 'a'
	str := "A árvore estava cheia de maçãs."
	analisaLetraA(str)

	// 3) Cálculo da variável SOMA
	imprimeSoma()

}
