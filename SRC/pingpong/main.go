package main

import (
	"fmt"
	"time"
)

func pin(ball chan <- int , action chan<- string )  {
    ball <-1
    action <-"Ping"
}
func pon(ball chan<- int , action chan<- string)  {
	ball <-2
	action <- "Png"
}
func refer(action <- chan string)  {
	for  {
	 fmt.Println("Action:",<-action)
	}
}
func game()  {
	ball :=make(chan int)
	action :=make(chan string)
    go refer(action)
	go pin(ball,action)
	for  {
		val:= <-ball
		switch val {
		case 1:
			go pon(ball,action)
		case 2:
			go pin(ball,action)

		}

	}
}
func main()  {
	go game()
	time.Sleep(10*time.Second)
}
