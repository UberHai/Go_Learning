package main

import (
	"fmt"
	"github.com/UberHai/Go_Learning/02_roman-numeral-server/romanNumerals"
	"html"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	//HTTP Package has methods for dealing with requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		urlPathElements := strings.Split(r.URL.Path, "/")
		//If request is GET with Correct Syntax
		if urlPathElements[1] == "roman_number" {
			number, _ :=
				strconv.Atoi(strings.TrimSpace(urlPathElements[2]))
			if number == 0 || number > 20 {
				//if resource is higher than 10 ( highest in data set )
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("404 - Roman Numeral Not Found"))
			} else {
				fmt.Fprintf(w, "%q",
					html.EscapeString(romanNumerals.Numerals[number]))
			}
		} else {
			//For all other requests tell the client it was a bad request
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - /roman_number/ only"))
		}
	})
	//Create a server and run it on port :8000

	s := &http.Server{
		Addr:           ":8000",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
