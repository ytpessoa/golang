package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} //array
	s1 := []int{1, 2, 3}  //slice

	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	//slice = pedaço de um array

	a2 := [5]int{1, 2, 3, 4, 5}
	s2 := a2[1:3] // indices 1 a 2
	fmt.Println(a2, s2)

	s3 := a2[:3]
	fmt.Println(s3)
}
