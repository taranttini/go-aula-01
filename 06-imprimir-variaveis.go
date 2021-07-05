package main

import "fmt"

func main() {

	fmt.Println("\n---------- variaveis 1 ----------")
	// criando variáveis simplificadas
	var texto1 = "texto"
	var numero_inteiro1 = 1
	var numero_real1 = 15.25
	var valor_logico_verdadeiro1 = true
	var valor_logico_falso1 = false
	var vetor1 = []int{1, 2, 3, 4}

	fmt.Println("texto =", texto1)
	fmt.Println("numero_inteiro =", numero_inteiro1)
	fmt.Println("numero_real =", numero_real1)
	fmt.Println("valor_logico_verdadeiro =", valor_logico_verdadeiro1)
	fmt.Println("valor_logico_falso =", valor_logico_falso1)
	fmt.Println("vetor =", vetor1)

	// as variáveis permitem a modificação do seu valor inicial
	// para um novo valor, mas respeitando o tipo inicial exemplo:
	// - texto só muda para outro texto
	// - número inteiro só muda para outro número inteiro
	// - número real só muda para outro número real
	texto1 = "texto_2"
	numero_inteiro1 = 2
	numero_real1 = 30.50
	valor_logico_verdadeiro1 = false
	valor_logico_falso1 = true
	vetor1 = []int{5, 6}

	fmt.Println("\n----- alterando os valores -----")
	fmt.Println("texto =", texto1)
	fmt.Println("numero_inteiro =", numero_inteiro1)
	fmt.Println("numero_real =", numero_real1)
	fmt.Println("valor_logico_verdadeiro =", valor_logico_verdadeiro1)
	fmt.Println("valor_logico_falso =", valor_logico_falso1)
	fmt.Println("vetor =", vetor1)

	// outras formas de declarar a variável

	fmt.Println("\n---------- variaveis 2 ----------")
	// criando variáveis em grupo
	var (
		texto2                   = "texto"
		numero_inteiro2          = 1
		numero_real2             = 15.25
		valor_logico_verdadeiro2 = true
		valor_logico_falso2      = false
		vetor2                   = []int{1, 2, 3, 4}
	)

	fmt.Println("texto =", texto2)
	fmt.Println("numero_inteiro =", numero_inteiro2)
	fmt.Println("numero_real =", numero_real2)
	fmt.Println("valor_logico_verdadeiro =", valor_logico_verdadeiro2)
	fmt.Println("valor_logico_falso =", valor_logico_falso2)
	fmt.Println("vetor =", vetor2)

	fmt.Println("\n---------- variaveis 3 ----------")
	// criando variáveis em linha, mais complexo se for muito longo
	var texto3, numero_inteiro3, numero_real3, valor_logico_verdadeiro3, valor_logico_falso3, vetor3 = "texto", 1, 15.25, true, false, []int{1, 2, 3, 4}

	fmt.Println("texto =", texto3)
	fmt.Println("numero_inteiro =", numero_inteiro3)
	fmt.Println("numero_real =", numero_real3)
	fmt.Println("valor_logico_verdadeiro =", valor_logico_verdadeiro3)
	fmt.Println("valor_logico_falso =", valor_logico_falso3)
	fmt.Println("vetor =", vetor3)

	fmt.Println("\n---------- variaveis 4 ----------")
	// criando variáveis declaracao curta
	texto4 := "texto"
	numero_inteiro4 := 1
	numero_real4 := 15.25
	valor_logico_verdadeiro4 := true
	valor_logico_falso4 := false
	vetor4 := []int{1, 2, 3, 4}

	fmt.Println("texto =", texto4)
	fmt.Println("numero_inteiro =", numero_inteiro4)
	fmt.Println("numero_real =", numero_real4)
	fmt.Println("valor_logico_verdadeiro =", valor_logico_verdadeiro4)
	fmt.Println("valor_logico_falso =", valor_logico_falso4)
	fmt.Println("vetor =", vetor4)
}
