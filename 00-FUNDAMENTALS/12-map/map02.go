package main

import "fmt"

func main() {
	// := declara e inicializa
	map_funcs_e_Salarios := map[string]float64{
		"Jose":   11000,
		"Pedro":  9000,
		"Ytallo": 10000,
	}

	map_funcs_e_Salarios["Gabriel"] = 8000
	delete(map_funcs_e_Salarios, "Inexistente") //nao acontece nada

	for nome, salario := range map_funcs_e_Salarios {
		fmt.Println(nome, salario)
	}

	for nome, _ := range map_funcs_e_Salarios {
		fmt.Println(nome)
	}
}
