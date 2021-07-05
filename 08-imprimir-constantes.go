package main

import "fmt"

func main() {

	fmt.Println("\n---------- constantes ----------")
	// segue o mesmo modelo das variáveis,
	// a diferença é que a variável pode receber um novo valor
	// já no caso da constante, seu valor não pode ser alterado
	const texto = "texto"
	const numero = 1

	fmt.Println("texto =", texto)
	fmt.Println("numero =", numero)

	// texto = "novo" // erro - cannot assign to texto (declared const)
	// numero = 2     // erro - cannot assign to numero (declared const)
}
