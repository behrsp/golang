package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//tipos inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é: ", reflect.TypeOf(32000))

	//sem sinal ( só positivos)...uint8 uint16, uint32, uint64
	var b byte = 255
	fmt.Println("O byte é:", reflect.TypeOf(b))

	//com sinal... int8, int16, int16, int32, int64
	i1 := math.MaxInt64
	fmt.Println("O valor maximo do int é ", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' //representa um mapeamento da tabela Unicode (int32)

	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	//numeros reais ( float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo literal é",reflect.TypeOf(49.99))

	//tipo boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println("O literal de bo é", reflect.TypeOf(true))
	fmt.Println(!bo)

	//tipo string
	s1 := "Ola meu nome é Robson"
	fmt.Println(s1+ "!")
	fmt.Println("O tamanho da string é",len(s1))

	//string com multiplas linhas
	s2 := `Olá
	meu 
	nome 
	é 
	Robson`
	fmt.Println("O tamanho da String é",len(s2))

	//char
	char := 'a'
	fmt.Println("o tipo de char é", reflect.TypeOf(char))

}
