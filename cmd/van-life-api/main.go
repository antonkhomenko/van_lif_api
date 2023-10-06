package main

import (
	"fmt"
	"log"
	"net/http"
	"van-life-api/internal/routes"
)

func main() {
	port := 8080
	laddr := fmt.Sprintf(":%d", port)
	router := routes.NewRouter()

	log.Printf("Server is runnig on port %s\n", laddr)
	log.Fatal(http.ListenAndServe(laddr, router))
}
