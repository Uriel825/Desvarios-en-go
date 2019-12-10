package main

import "fmt"

type Platzi struct {
	name string
	slug string          //estrcuctura como c++
	skill [] string
}
type PlatzyCoures struct {
	Platzi
}

func (p Platzi)  subcribe (name string){
	fmt.Printf("%s","se ha registrado al curso ",'\n' , name,p.name)
}//solo si se llama de una ubicacion solo se puede hacer un acceso a a la vez

//Interfaces
type Platzy interface {
	Platzi
}

func callSub(platzi Platzi)  {
platzi.subcribe("HO")
}



func main ()  {

	platzi := Platzi{
		name:  "GO",
		slug:  "GO.com",
		skill: []string{"G","o"},
	}
	platzi.subcribe("COCO")
	platzi1 := new(Platzi) //se puede instanciar
	platzi1.name="c++"
	fmt.Print(platzi)
 }
//revisar