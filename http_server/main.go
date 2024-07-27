package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = ":9000"

func main() {
	port := PORT
	server := NewPlayerServer(NewInMemoryPlayerStore())
	log.Fatal(http.ListenAndServe(port, server))
	fmt.Printf("listening on port %q", port)
}
