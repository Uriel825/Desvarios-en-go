package main

import (
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request)  {
io.WriteString(w,"Holi") //devuelve una string
}
func main()  {
	http.HandleFunc("/",handler)
	http.ListenAndServe(";8000",nil)
}
