package print

import "fmt"

func Print() string { // mayusculas es public
	var algo string
	fmt.Print("Hello world","Ingresa algo")
	fmt.Scanf("%s",&algo)
    fmt.Print(algo)
	return ""
}