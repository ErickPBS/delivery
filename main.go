package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

func carrinhoHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "carrinho.html", nil)
}

func main() {
	// Servir arquivos estáticos (CSS, JS, imagens)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Rotas
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/carrinho", carrinhoHandler)

	log.Println("Servidor rodando em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
