package main

import (
	"fmt"
)

// Aqui está um exemplo de declaração de uma constante em Go
const CPF = "123.432.678-32"

// Constantes tipadas são declaradas com um tipo definido
const TELEFONE = 11999999999

// Várias constantes podem ser agrupadas em um bloco para facilitar a leitura
const (
	A int = 1
	B float32 = 3.14
	C = "H1!"
)

func main () {
	fmt.Println(CPF)
	fmt.Println("============")
	fmt.Println(TELEFONE)
	fmt.Println("============")
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}