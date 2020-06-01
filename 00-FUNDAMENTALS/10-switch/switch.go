package main

import "fmt"

func notaParaConceito(n float64) string {
	var nota = int(n)
	switch nota {
	case 10:
		return "10"
	case 9, 8:
		return "A"
	case 7, 6:
		return "B"
	default:
		return "Nota Invalida"
	}
}

func main() {

	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(-1))
}
