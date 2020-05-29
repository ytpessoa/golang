package main

import "fmt"

func main() {
	fmt.Println("no final quebra a linha")
	fmt.Print("mesma linha")
	fmt.Println("")

	x := 3.141516
	xs := fmt.Sprint(x)
	fmt.Println("O valor de x e " + xs)
	fmt.Println("O valor de x e", x)

	fmt.Printf("O valor de x e %.2f", x)

}
