package op_cadena

import (
	"fmt"
	"strings"
) // se importa la libreria string

func str()  {
	var text   = "Hola,como estas,buenas, noches"
    fmt.Println(strings.ToUpper(text)) //mayusculsas
	fmt.Println(strings.ToLower(text)) //minusculas
	fmt.Println(strings.Replace(text , "Hola","Hello" ,-1)) //todos los de old
	fmt.Println(strings.Split(text," ")) //separar por " " comas,espacios o puntos
}
func main()  {
	str()
}