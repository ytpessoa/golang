package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	x := 3
	fmt.Println("O tipo de x eh", reflect.TypeOf(x))

	// sem sinal = apenas positivos
	var b byte = 255
	fmt.Println("O byte e", reflect.TypeOf(b)) //uint8

	i1 := math.MaxInt64
	fmt.Println("O valor maximo do int eh", i1)

	s1 := "Olá meu nome é Ytallo"
	fmt.Println(s1 + "!")
	fmt.Println(len(s1))
	s2 := `Ola
	meu nome
	e Gabriel`
	fmt.Println(s2)
}
