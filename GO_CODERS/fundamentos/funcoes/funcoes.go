package main 

import "fmt"


func somar(a int, b int) int {
//passar por paramtro tudo que vc quer passar de entrada
//na funcao vai, nome do parametro e tipo
// e no final o tipo de retorno

// como temos a palavra reservada int
// tem que chamar o return quando ao tiver nao precisa exemplo abaixo

	return a + b
}

func imprimir(valor int){
// neste caso nao tem o int depois da funcao, entao nao precisa chamar o int

	fmt.Println(valor)
}

