package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var guess int
	var limi, limf int = 1, 9999

	fmt.Println("Adivinhar Numero")
	limf = rand.Intn(limf-limi) + limi

	fmt.Printf("Adivinhe um numero entre %d e %d: ", limi, limf)
	fmt.Scanf("%d", &guess)

}
