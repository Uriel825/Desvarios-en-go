package main

import "fmt"
//Las constantes se crean fuera de la funcion
const hello string = "Hola"
func main() {
	var name string
	lastName := "pica"// tambien se puede declarar una variable asi
	var number = 100               // sabe en automatico que es un numero
    var (         //tambien se puede usar esto para multiples
    	a=1
    	b=2
    	c=3
	)
	fmt.Println("Dime tu nombre")
	fmt.Scanf("%s",&name)
	fmt.Println("Dime tu apellido")
	fmt.Scanf("%s",lastName)
	fmt.Println(hello,name,lastName)
	fmt.Println("Selecciona un numero",a,",",b,",",c,",","Asi hasta el",number)
}