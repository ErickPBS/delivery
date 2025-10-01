package rotas

import (
	"modulo/controle"
	"net/http"
)

func Rotas() {

	http.HandleFunc("/", controle.Index)
	http.HandleFunc("/login", controle.Login)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

}
