package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Uso: conversor <valores> <unidade>")
		os.Exit(1)
	}

	unidadeOrigem := os.Args[len(os.Args)-1]
	valoresOrigem := os.Args[1 : len(os.Args)-1]

	var unidadeDestino string

	if unidadeOrigem == "c" {
		unidadeDestino = "f"
	} else if unidadeOrigem == "f" {
		unidadeDestino = "c"
	} else {
		fmt.Printf("%s não é uma unidade conhecida!\n", unidadeOrigem)
		os.Exit(1)
	}

	for i, v := range valoresOrigem {
		valorOrigem, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Printf("O valor %s não posição %d não é um numero válido!\n", v, i)
			os.Exit(1)
		}

		var valorDestino float64

		if unidadeOrigem == "c" {
			valorDestino = valorOrigem*1.8 + 32
		} else {
			valorDestino = (valorOrigem - 32) * 5 / 9
		}

		fmt.Printf("%.2f º%s = %.2f º%s\n",
			valorOrigem, unidadeOrigem,
			valorDestino, unidadeDestino)

	}
}
