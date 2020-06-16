package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var guess, generated int
	var limi, limf int = 1, 9999
	var lives = 3

	fmt.Printf("Adivinhar Numero - Vidas: %d\n", lives)

	rand.Seed(time.Now().UnixNano())
	limf = rand.Intn(limf-limi) + limi

	generated = rand.Intn(limf-limi) + limi

	fmt.Printf("Adivinhe um numero entre %d e %d: ", limi, limf)
	fmt.Scanf("%d", &guess)
	fmt.Printf("Gerado: %d\n", generated)
}
