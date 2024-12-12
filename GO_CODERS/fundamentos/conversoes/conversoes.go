package main

import (
	"fmt"
	"strconv"
)

func main() {

	x := 2.4
	y := 2
	fmt.Println(x / float64(y))


	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	//cuidado ao contatenar
	fmt.Println("Teste" + string(123)) //neste caso o valor 123 representa o { da tabela ASCII

	//int  para string
	fmt.Println("Teste " + strconv.Itoa(123))

	// String para int
	num, _:= strconv.Atoi("123")
	fmt.Println(num - 122)

	b, _:= strconv.ParseBool("false")
	if b {
		fmt.Println("Verdadeiro")
	}else{
		fmt.Println("Falsoooo")
	}


}
