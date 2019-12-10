package main

import "fmt"


func main() {
	selector:=0
		fmt.Println("Calculadora:", "\n", "Suma presiona (1)", "\n", "Resta presiona (2)", "\n",
		"Multiplicacion presiona (3)", "\n", "Division presiona (4)")
		fmt.Scanf("%d", &selector)

	switch selector{
	case 1:
		fmt.Print("El resultado es:", sum())
		break
	case 2:
		fmt.Print("El resultado es:", res())
		break
	case 3:
		fmt.Print("El resultado es:", mult())
		break
	case 4:
		fmt.Print("El resultado es:", div())
		break

		default:
			fmt.Print("Error opcion invalida")
			break
		}

}
func sum()float32 {
	var suma ,a,b float32
	fmt.Println("Ingresa el primer numero")
	fmt.Scanf("%f",&a)
	fmt.Println("Ingresa el segundo numero")
	fmt.Scanf("%f",&b)
	suma = (a+b)
	return suma
}

func res()float32  {
	var rest,a,b float32
	fmt.Println("Ingresa el primer numero")
	fmt.Scanf("%f",&a)
	fmt.Println("Ingresa el segundo numero")
	fmt.Scanf("%f",&b)
	rest = (a-b)
	return rest
}

func mult()float32  {
	var mult,a,b float32                    //int=32bits
	fmt.Println("Ingresa el primer numero")
	fmt.Scanf("%f",&a)
	fmt.Println("Ingresa el segundo numero")
	fmt.Scanf("%f",&b)
	mult = (a*b)
	return mult
}

func div()float32  {
	var divi,a,b float32
	fmt.Println("Ingresa el primer numero")
	fmt.Scanf("%d",&a)
	fmt.Println("Ingresa el segundo numero")
	fmt.Scanf("%d",&b)
	divi = (a/b)
	return divi
}