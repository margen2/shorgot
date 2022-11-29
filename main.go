package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/margen2/shorgot/src/router"

	"github.com/margen2/shorgot/src/config"
)

func main() {
	config.Load()
	fmt.Printf("Running API on port :%d", config.Port)
	r := router.GenerateRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", config.Port), r))
}
