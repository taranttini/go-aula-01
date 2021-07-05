package main

import "fmt"

func main() {

	texto := "texto"           // string
	caracter := 'C'            // rune / int32
	numero_inteiro := 1        // int
	numero_real := 15.25       // float64
	valor_logico := true       // bool
	vetor := []int{1, 2, 3, 4} // []int
	complexo := 6i             //complex

	fmt.Printf("texto, tipo = %T \n", texto)
	fmt.Printf("caracter, tipo = %T \n", caracter)
	fmt.Printf("numero_inteiro, tipo = %T \n", numero_inteiro)
	fmt.Printf("numero_real, tipo = %T \n", numero_real)
	fmt.Printf("valor_logico, tipo = %T \n", valor_logico)
	fmt.Printf("vetor, tipo = %T \n", vetor)
	fmt.Printf("complexo, tipo = %T \n", complexo)

	// não é permitido somar diferentes tipos
	/*
		exemplo1 := texto + numero_inteiro // erro (mismatched types string and int)
		exemplo2 := numero_inteiro + numero_real // erro (mismatched types int and float64)
		exemplo3 := numero_inteiro + valor_logico // erro (mismatched types int and float64)
	*/
}
