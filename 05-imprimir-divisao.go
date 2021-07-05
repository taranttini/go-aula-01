package main

import "fmt"

func main() {
	fmt.Println("1 ÷ 1 =", 1/1)
	fmt.Println("1 ÷ 2 =", 1/2)          // não vai imprimir fracionado
	fmt.Println("1 ÷ 2 =", 1/2.0)        // ou fmt.Println("1 ÷ 2 = ", 1.0/2)
	fmt.Println("1 ÷ 2 =", float64(1)/2) // ou fmt.Println("1 ÷ 2 = ", 1/float64(2))
	fmt.Println("2 ÷ 1 =", 2/1)
}
