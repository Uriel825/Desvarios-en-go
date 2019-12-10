package main


import "fmt"
func poin(){
	a := 100
	var  b *int
	b=&a
	fmt.Print(b)

}
func main()  {
	poin()
}