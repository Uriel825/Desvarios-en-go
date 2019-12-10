package main

import "fmt"

func ifTest()  {
	var number = 0
	var number2  = 0
	fmt.Print("Ingresa un numero")

	fmt.Scanf("%d",&number) // siempre va el apuntador a la variable
	if number % 2 == 0 {
		fmt.Print("es par")
	}else{
		fmt.Print("es impar")
	}
	if number2 := 3; number2 ==3{
		fmt.Println("Entro al condicional")
	}
}
func main()  {
	ifTest()
}