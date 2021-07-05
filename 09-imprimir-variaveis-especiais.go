package main

import "fmt"

func main() {

	texto1 := "texto" // valor
	texto2 := &texto1 // endereçamento de memória do texto1
	texto3 := *texto2 // ponteiro que contem o valor do endereço da memória de texto2, (o resultado é o mesmo de texto1)

	fmt.Println("texto1 =", texto1)
	fmt.Println("&texto2 =", texto2)
	fmt.Println("*texto3 =", texto3)
}
