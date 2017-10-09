// execService.go
package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()
	//Mapping to methods is possible with httprouter
	//Get request to /static/file.txt sends back file
	router.ServeFiles("/static/*filepath", http.Dir("./static/"))
	log.Fatal(http.ListenAndServe(":8000", router))
}
