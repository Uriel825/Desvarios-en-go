package print

import "fmt"


func print()  {       //la primera de la letra de la funcion indica si es public o private
                      // minuscula private , mayuscula public
	var name string // Declaracion de la variable
	fmt.Print("Cual es tu nombre?") // Imprime a pantalla
	fmt.Scanf("%s" , &name) // %s: dicta que la variable a tomar es un string , &v(var) apunta a la variable de guardado
	fmt.Printf("Hola,%s",name) //imprime diciendo el tipo de variable y la cvariabe

}