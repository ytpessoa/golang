//funcoes variativas: recebem parametros variaveis

package main

import "fmt"

func media(numeros ...float64) float64 {
	//numeros entra como array
	total := 0.0
	for _, num := range numeros {
		total += num
	}

	if total > 0.0 {
		return total / float64(len(numeros))
	} else {
		return 0.0
	}

}

func main() {
	fmt.Printf("media eh %f", media())
}
