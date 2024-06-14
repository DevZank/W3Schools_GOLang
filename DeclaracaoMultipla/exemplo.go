package main

import (
	"fmt"
)

func main() {
	// No Go, é possível declarar múltiplas variáveis ​​na mesma linha.
	var a, b, c, d int = 1, 3, 5, 7

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	// Se um tipo de palavra-chave não for especificado, você poderá declarar diferentes tipos de variáveis ​​na mesma linha.
	var e, f = 6, "Hello"
	g, h := 7, "World!"

	fmt.Println("=============")
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)

	// Múltiplas declarações de variáveis ​​também podem ser agrupadas em um bloco para maior legibilidade.
	var (
		i int
		j int = 1
		k string = "Hello"
	)

	fmt.Println("=============")
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
}