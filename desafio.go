package main

import "fmt"

//(FORMULA DA CONVERSAO) DICA C = K - 273
//O VALOR DO PONTO DE EBULIÇÃO DA AGUA EM KELVIN É DE 373°
const tempEbulicaoK = 373.0

func main() {

	tempK := tempEbulicaoK
	tempC := (tempK - 273.0)

	fmt.Printf("A temperatura de ebulição em °K = %g (%T) e a temperatura de ebulição em °C é = %g (%T)", tempK, tempK, tempC, tempC)
	
	// O PROGRAMA IRÁ IMPRIMIR A TEMPERATURA EM °K E EM °C

}
