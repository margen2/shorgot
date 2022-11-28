package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/margen2/shorgot/src/router"

	"github.com/margen2/shorgot/src/config"
)

//func init() {
//	key := make([]byte, 64)
//	if _, err := rand.Read(key); err != nil {
//		log.Fatal(err)
//	}
//	stringBase64 := base64.StdEncoding.EncodeToString(key)
//	fmt.Println(stringBase64)
//}

func main() {
	config.Load()
	fmt.Printf("Running API on %s", config.Port)
	r := router.GenerateRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s", config.Port), r))

}
