package main

import "fmt"

// func nameFunc (paramatro1 tipo, parametro2 tipo) (tipoRetorno) {

// }

func f1() {
	fmt.Println("F1")
}
func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s, %s\n", p1, p2)
}

func f3() string {
	return "F3"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

//multiplos retornos
func f5() (string, string) {
	return "Retorno da funcao 1", "Retorno da funcao 2"
}

func main() {
	f1()
	f2("param1", "param2")

	result3, result4 := f3(), f4("result4.1", "result4.2")
	fmt.Println(result3)
	fmt.Println(result4)

	retorno1, retorno2 := f5()
	fmt.Println(retorno1)
	fmt.Println(retorno2)

}
