package main

import (
	"errors"
	"fmt"
)

func Sum(a interface{}, b interface{}) (int,error) {

	switch a.(type) {
	case string:
		return 0,errors.New("El valor a es un string")
	}
	switch b.(type){
	case string:
		return 0 ,errors.New("El valor a es un string")
	}

	return a.(int)+b.(int) ,nil
}

func main () {
	fmt.Print(Sum(50 , "a"))

}