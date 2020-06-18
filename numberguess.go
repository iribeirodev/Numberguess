package main

import (
	"C"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var limi, limf int = 1, 9999
var lives int = 5
var guess, generated int
var acertou bool = false

func obterPalpite() {
	fmt.Printf("Adivinhe um numero entre %d e %d: ", limi, limf)
	fmt.Scanf("%d", &guess)
}

func gerarResposta() {
	generated = rand.Intn(limf-limi) + limi
	for generated%10 != 0 {
		generated = rand.Intn(limf-limi) + limi
	}

	//fmt.Printf("O GERADO FOI %d\n", generated)
}

func controlarVidas() {
	if guess != generated {
		configurarLimite()
		lives--
	} else {
		acertou = true
	}
}

func configurarLimite() {
	if guess > generated {
		limf = guess
	} else if guess < generated {
		limi = guess
	}
}

func delimitar() {
	var delimi = strings.Repeat("~", 80)
	fmt.Println(delimi)
}

func exibirRelatorio() {
	clearScreen()
	delimitar()
	if !acertou {
		fmt.Println("Acabaram-se as vidas, jogo terminado.")
		fmt.Printf("O número correto era %d\n", generated)
		fmt.Println("Melhor sorte da próxima vez e Muito Obrigado por jogar conosco.")
	} else {
		fmt.Printf("PARABÉNS, VOCÊ ACERTOU !!!!  RESPOSTA CORRETA: %d.\n", generated)
	}
	delimitar()
}

func clearScreen() {
	print("\033[H\033[2J")
}

func configurarApp() {
	// Randomizacao
	rand.Seed(time.Now().UnixNano())
	// Limpando a tela
	clearScreen()
}

func exibirScore() {
	fmt.Printf("Adivinhar Numero - Você ainda tem %d Vida(s).\n", lives)
}

func main() {
	configurarApp()
	gerarResposta()

	for lives > 0 {
		exibirScore()
		obterPalpite()
		controlarVidas()
		if acertou {
			break
		}
	}

	exibirRelatorio()
}
