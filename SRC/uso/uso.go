package main

import "fmt"

func main()  {
	var name string
	fmt.Print("Como te llamas")
    fmt.Scanf("%s" ,&name)
	fmt.Printf("Que onda %s como esta? \n",name )
}