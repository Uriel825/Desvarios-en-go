package main

import (
	"fmt"
	"time"
)

func helloGO(index int)  {
	fmt.Println("HOla soy un go#",index)
}

func forGo(n int)  {
	for i:=0;i<n;i++  {
		go helloGO(i) //ejecuta en su propio hilo
	}
}

func main()  {
go forGo(500)
go forGo(700)
time.Sleep(10000*time.Millisecond) //time para evitar que la funcion main acabe primero que las go routines
}