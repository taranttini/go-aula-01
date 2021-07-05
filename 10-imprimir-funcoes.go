package main

import (
	"errors"
	"fmt"
)

func main() {

	mensagem()

	fmt.Printf("1 + 2 = %v \n", soma(1, 2))
	fmt.Printf("2 - 1 = %v \n", subtracao(2, 1))
	fmt.Printf("2 x 4 = %v \n", multiplicacao(2, 4))
	fmt.Printf("9 ÷ 3 = %v \n", divisao(9, 3))

	// função chamando outras funções
	fmt.Println(divisao(multiplicacao(3, 3), soma(1, 2)))

	// validando funcao com retorno de erro
	erro := nao_pode_ser_maior_que_10(20)
	if erro != nil {
		fmt.Println(erro.Error())
	}
}

// função sem parâmetro e sem retorno
func mensagem() {
	fmt.Println("realizando os seguintes cálculos:")
}

// função com parâmetro e com retorno
func soma(valor_x int, valor_y int) int {
	return valor_x + valor_y
}

// função com parâmetro e com retorno
func subtracao(valor_x int, valor_y int) int {
	return valor_x - valor_y
}

// função com parâmetro e com retorno
func multiplicacao(valor_x int, valor_y int) int {
	return valor_x * valor_y
}

// função com parâmetro e com retorno
func divisao(valor_x int, valor_y int) int {
	return valor_x / valor_y
}

// função que realiza uma lógica e gera erro caso não atenda as necessidades
func nao_pode_ser_maior_que_10(valor int) error {
	if valor > 10 {
		return errors.New("valor não pode ser maior que 10")
	}
	return nil
}
