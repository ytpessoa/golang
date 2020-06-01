package main

import "fmt"

func main() {

	i := 1
	var p *int = nil

	p = &i
	*p++

	fmt.Println(p, &i)
	//0xc0000160d0  0xc0000160d0
	fmt.Println(*p, i)
	// 2 2
}
