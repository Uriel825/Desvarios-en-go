package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Album struct {
	userId  int `json:"userId"` //etiquetas para saber que tipo de datos debe manejar 
	id int    `json:"id"`        //estrucutura para leer la informacion y construirla
	title string `json:"title"`
}

func main() {
	var client = &http.Client{Timeout: 10 * time.Second}
	url := "https://jsonplaceholder.typicode.com/albums"
	resp, err := client.Get(url)
	if err != nil {
		panic(err.Error())
	}
	var album []Album       //debe de ir en array por que sino no hace match con los datos que se requieren  usar
    err = json.NewDecoder(resp.Body).Decode(&album) //para la decodificacion de los datos
	if err != nil {
		panic(err.Error())

	}
	fmt.Println(album[1]) // imprime todos
}
nombre = "PLatzi"
const name  = 