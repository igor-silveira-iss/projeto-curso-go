// Package utils centraliza as utilidades da aplicação web
package utils

import (
	"net/http"
	"text/template"
)

var templates *template.Template

// CarregarTemplates insere os templates html na variável templates
func CarregarTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

// ExecutarTemplate renderiza uma página html na tela
func ExecutarTemplate(w http.ResponseWriter, template string, dados any) {
	templates.ExecuteTemplate(w, template, dados)
}
