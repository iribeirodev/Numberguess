package main

import (
	"fmt"
	"math/rand"
	"time"
  "strings"
)

var limi, limf int = 1, 9999
var lives int = 5
var guess, generated int

func obterPalpite() {
	fmt.Printf("Adivinhe um numero entre %d e %d: ", limi, limf)
	fmt.Scanf("%d", &guess)
}

func gerarResposta() {
	generated = rand.Intn(limf-limi) + limi
  for generated % 10 != 0 {
    fmt.Printf("gerando numero: %d\n", generated)    
    generated = rand.Intn(limf-limi) + limi
  }

  fmt.Printf("O GERADO FOI %d\n", generated)
}

func controlarVidas() {
	if guess != generated {
		configurarLimite()
		lives--
	} else {
    fmt.Printf("PARABÉNS, VOCÊ ACERTOU !!!!  RESPOSTA CORRETA: %d.\n", generated)
    fmt.Println("VOCÊ É O REI DOS CHUTES. GOSTARIA DE TENTAR MAIS UMA RODADA?")
  }
}

func configurarLimite() {
  if guess > generated {
    limf = guess
  } else if guess < generated {
    limi = guess
  }
}

func exibirRelatorio() {
  delimi := strings.Repeat("~", 80)
  clearScreen()

  fmt.Println(delimi)

  fmt.Println("Acabaram-se as vidas, jogo terminado.")
  fmt.Printf("O número correto era %d\n", generated)
  fmt.Println("Melhor sorte da próxima vez e Muito Obrigado por jogar conosco.")

  fmt.Println(delimi)
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
  fmt.Printf("Adivinhar Numero - Vidas: %d\n", lives)
}

func main() {
  configurarApp()
	gerarResposta()

	for lives > 0 {
    exibirScore()
		obterPalpite()
		controlarVidas()
	}

  exibirRelatorio()
}
