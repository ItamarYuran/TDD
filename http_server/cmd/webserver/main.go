package main

import (
	"log"
	"net/http"

	poker "github.com/ItamarYuran/TDD/http_server"
)

const dbFileName = "../../game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	server := poker.NewPlayerServer(store)

	if err := http.ListenAndServe(":9000", server); err != nil {
		log.Fatalf("could not listen on port 9000 %v", err)
	}
}
