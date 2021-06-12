package main

import "fmt"

func main() {
	var variavel1 string = "variavel1"
	variavel2 := "variavel2"
	fmt.Println(variavel1, variavel2)
	// Outros tipos

	var (
		variavel3 string = "variavel3"
		variavel4 string = "variavel4"
	)

	variavel5, variavel6 := "variavel5", "variavel6"

	fmt.Println(variavel3, variavel4)
	// Como trocar o valor de uma variável entre si
	fmt.Println(variavel5, variavel6)
	variavel5, variavel6 = "variavel6", "variavel5"
	fmt.Println(variavel5, variavel6)

	//Como criar constantes

	const constante1 string = "constante1"
	fmt.Println(constante1)
}
