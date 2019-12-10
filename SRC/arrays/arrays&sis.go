package main

import "fmt"

func getArray (){
	var arr1  [2] string  // de largo fijo y tipo
    arr2 := [3]  int{1,2,3} // tambien se pueden llenar de esta forma
	arr1[0] = "Hola"
    arr1[1]="soy un array"
    fmt.Print(arr1,arr2)
}
func getSlice (){
	var slice []string
	slice = append(slice,"hola","que","hay") // para llenar un slice se usa append
	fmt.Print(slice)
}

func main()  {
	getArray() // imprimir
    getSlice()
}