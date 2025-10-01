package main

import (
	"fmt"
	"modulo/rotas"
	"net/http"
)

func main() {
	rotas.Rotas()
	fmt.Println("servidor rodando na porta 9090")
	http.ListenAndServe(":9090", nil)
}
