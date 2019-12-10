package main

import "fmt"

func main()  {
	ch:= make(chan string) // crear el canal
	go func() { //funcion anonima
		defer close(ch) //cerrar el canal y evitar que la variable se sobreescriba
		ch <- "Hola Chanel"
	}()
	fmt.Print(<-ch) //<- ch obtiene el valor de la variable
	// ch <- envia datos
ch1 := make(chan int)
go func() {
	defer close(ch1)
	ch1 <- 1
	ch1 <-2
}()
	for n := range ch1{
		fmt.Printf("%d",n)
	}
ch2:= make(chan int ,2) //canal con bufer se especifica la cantidad maxima
ch2 <- 1
ch2<-2
fmt.Print(<-ch2)
fmt.Print(<-ch2)
ch2<-3
fmt.Print(ch2)
}
