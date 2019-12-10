package main

import (
	"fmt"
)
//mapas estilo diccionarios
func maps(name string) int {
	mapTest :=  map[string]int{   // make (map[string]int)
		 "Juan": 28,
		"Pepe": 32,
	}
    mapTest["llave 1 "] =1
    mapTest["llave 2 "] =2
   delete(mapTest,"llave 1")
    return mapTest[name]
}

func main()  {
	fmt.Print(maps("Juan"))
}