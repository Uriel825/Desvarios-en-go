package main
// en go solo hay un ciclo for

import "fmt"

func forT()  {
	for i:= 1; i<=5; i++ {
		fmt.Println("FOR", i)
	}
}
func forU()  {
	c:=100     //otra manera del loop for

	for; c>0; {
	c-=10
		println(c)
	}
}
func forI()  { // loop infinito
	s:=1000
	for  {
		s-=1
		if s == 0 {
			fmt.Println("termino")
		break
		}
	}
}
func main()  {
	forT()
	forU()
	forI()
}