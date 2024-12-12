package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println(" Nova")
	fmt.Println("linha.")

	x := 3.141516

	//fmt.Println("O valor de x é " + x)   isso nao é permitido em go


	// a formas corretas
	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é: ", x)

	//ou
	fmt.Printf("O valor de x é %f", x)
	fmt.Printf("O valor de x é %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "Opa"

	fmt.Printf("\n%d %f %.1f %t %s",a,b,b,c,d)
}
