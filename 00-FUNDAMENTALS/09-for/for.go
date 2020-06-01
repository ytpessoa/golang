//não tem "while" em go
package main

import "fmt"

func main() {

	i := 1

	for i <= 10 { //while
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

}
