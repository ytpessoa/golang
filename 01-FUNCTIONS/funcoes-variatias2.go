//funcoes variativas: recebem parametros variaveis

package main

import "fmt"

func printAprovados(aprovados ...string) {
	for _, aluno := range aprovados {
		fmt.Println(aluno)
	}
}
func main() {
	aprovados := []string{"ytallo", "gabriel", "pessoa", "leda"} //slice=nao definimos tamanho
	printAprovados(aprovados...)
}
