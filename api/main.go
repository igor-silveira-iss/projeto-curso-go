package main

import (
	"fmt"
	"log"
	"net/http"

	"api/src/router"
)

func main() {
	fmt.Println("Rodando API")
	r := router.Gerar()
	log.Fatal(http.ListenAndServe(":5000", r))
}
