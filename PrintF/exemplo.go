package main

import (
	"fmt"
)

func main() {
	// A função Printf() primeiro formata seu argumento com base no verbo de formatação fornecido e depois os imprime.

	var i string = "Hello"
	var j int = 15

	// %v é usado para imprimir o valor dos argumentos
	// %T é usado para imprimir o tipo dos argumentos
	fmt.Printf("i has value: %v and type: %T \n", i, i)
	fmt.Printf("j has value: %v and type: %T \n", j, j)

	var a = 15.5
	var b = "Hello World!"

	fmt.Printf("=========== \n")
	fmt.Printf("%v\n", a)
	fmt.Printf("%#v\n", a)
	fmt.Printf("%v%%\n", a)
	fmt.Printf("%T\n", a)
	fmt.Printf("=========== \n")
	fmt.Printf("%v\n", b)
	fmt.Printf("%#v\n", b)
	fmt.Printf("%T\n", b)

	var c = 15

	// %b                               Base 2
	// %d                               Base 10
	// %+d                             Base 10 e sempre mostra sinal
	// %o                               Base 8
	// %O                              Base 8, com 0o inicial
	//%x                               Base 16, minúscula
	//%X                               Base 16, maiúscula
	// %#x                           Base 16, com 0x inicial
	// %4d                          Bloco com espaços (largura 4, justificado à direita)
	// %-4d                        Bloco com espaços (largura 4, justificado à esquerda)
	// %04d                       Bloco com zeros (largura 4

	fmt.Printf("=========== \n")
	fmt.Printf("%b\n", c)
	fmt.Printf("%d\n", c)
	fmt.Printf("%+d\n", c)
	fmt.Printf("%o\n", c)
	fmt.Printf("%O\n", c)
	fmt.Printf("%x\n", c)
	fmt.Printf("%X\n", c)
	fmt.Printf("%#x\n", c)
	fmt.Printf("%4d\n", c)
	fmt.Printf("%-4d\n", c)
	fmt.Printf("%04d\n", c)

	var d = "Hello"

	// %s            Imprime o valor como string simples
	// %q           Imprime o valor como uma string entre aspas duplas
	// %8s         Imprime o valor como string simples (largura 8, justificado à direita)
	// %-8s        Imprime o valor como string simples (largura 8, justificado à esquerda)
	// %x           Imprime o valor como dump hexadecimal de valores de bytes
	//% x           Imprime o valor como dump hexadecimal com espaços

	fmt.Printf("=========== \n")
	fmt.Printf("%s\n", d)
	fmt.Printf("%q\n", d)
	fmt.Printf("%8s\n", d)
	fmt.Printf("%-8s\n", d)
	fmt.Printf("%x\n", d)
	fmt.Printf("% x\n", d)

	var e = true
	var f = false

	// %t	Valor do operador booleano no formato verdadeiro ou falso (o mesmo que usar %v)

	fmt.Printf("=========== \n")
	fmt.Printf("%t\n", e)
	fmt.Printf("%t\n", f)

	var g = 3.141

	// %e Notação científica com 'e' como expoente
	// %f Ponto decimal, sem expoente
	// %.2f Largura padrão, precisão 2
	// %6.2f Largura 6, precisão 2
	// %g Expoente conforme necessário, apenas dígitos necessários

	fmt.Printf("=========== \n")
	fmt.Printf("%e\n", g)
	fmt.Printf("%f\n", g)
	fmt.Printf("%.2f\n", g)
	fmt.Printf("%6.2f\n", g)
	fmt.Printf("%g\n", g)
}