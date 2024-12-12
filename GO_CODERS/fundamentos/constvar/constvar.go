package main

import (
	"fmt"
	m "math"
)

func main() {

	const PI float64 = 3.1415 //criar constante
	var raio = 3.2            // tipo (float64)

	//forma reduzida igual ao advpl
	area := PI * m.Pow(raio, 2)
	fmt.Println("Area da circunferencia é: ", area)

	//outra forma para declarar constante
	const (
		a = 1
		b = 2
	)
	// mesma forma para definir variaveis
	var (
		c = 3
		d = 4
	)

	fmt.Println(a,b,c,d)

	// outra forma de declarar
	// no caso true vai para e , e false para f
	var e, f bool = true, false 

	fmt.Println(e,f)

	g,h,i := 2, false,"epa!"

	fmt.Println(g,h,i)

}
