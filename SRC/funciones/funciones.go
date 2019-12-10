package main

import "fmt"

func getNAme()string  {
	var name string
	fmt.Println("Cual es tu nombre")
	fmt.Scanf("%s",&name)
	return name
}
func numbers ()(int,int,int) {
	return 1, 2, 3
	}

func main() {
	user := getNAme()
	a,b,c := numbers()
	fmt.Println("hola",user)
    fmt.Println(a,b,c)
}
