package main

import (
	"fmt"
	"log"
	"net/http"

	"api/src/config"
	"api/src/router"
)

func main() {
	config.Carregar()
	// fmt.Println(config.Porta)
	// fmt.Println(config.StringConexaoBanco)
	fmt.Println("Rodando API")
	r := router.Gerar()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
