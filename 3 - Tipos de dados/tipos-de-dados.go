package main

import (
	"errors"
	"fmt"
)

func main() {
	// Existem 4 tipos de números inteiros no GO: int8,int16, int32 e int64 (a diferença está em quanto cada um suporta em bits)
	var numero1 int8 = 10
	fmt.Println(numero1)

	var numero2 int16 = 100
	fmt.Println(numero2)

	var numero3 int32 = 1000
	fmt.Println(numero3)

	var numero4 int64 = 1000000
	fmt.Println(numero4)

	// Existe também o int sozinho, que é quando você não especifica. Ele utiliza a arquitetura do seu computador como base

	var numero5 int = 100000000
	fmt.Println(numero5)

	// Inferência de tipo (Outro tipo de realizar variável):

	numero6 := 1000000000
	fmt.Println(numero6)

	// O GO também suporta números negativos

	numero7 := -1000000000
	fmt.Println(numero7)

	// UINT serve para representar um int sem sinal

	var numero8 uint64 = 44541
	fmt.Println(numero8)

	// Alguns ints possuem ALIAS que são como apelidos:

	// Alias para int32
	var numero9 rune = 12345
	fmt.Println(numero9)

	// Alias para uint8
	var numero10 byte = 12
	fmt.Println(numero10)

	// Números reais utilizam float32 e float64

	var numero11 float64 = 19.11
	fmt.Println(numero11)

	numero12 := 19.19
	fmt.Println(numero12)

	// Strings (Sempre utilizar aspas duplas):

	var str string = "Texto"
	fmt.Println(str)

	// Quando utilizado aspas simples, é convertido para o número da tabela ASC

	char := 'A'
	fmt.Println(char)

	// Valor 0

	var str2 string
	fmt.Println(str2)

	var numero13 int64
	fmt.Println(numero13)

	// Tipo de dados booleano

	var booleano1 bool = true
	fmt.Println(booleano1)

	// Se deixar sem valor na variável, o valor 0 do boolean é false

	var booleano2 bool
	fmt.Println(booleano2)

	// Tipo de dado <erro>

	var erro error
	fmt.Println(erro)

	// Caso queira dar um nome ao erro, é necessário utilizar o pacote errors

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}
