package main

import (
	"log"
	"os"
)

func main() {
	cfg := config{
		addr: ":8080",
		db:   dbConfig{},
	}

	api := application{
		config: cfg,
	}

	h := api.mount()
	if err := api.run(h); err != nil {
		log.Printf("server has failed to star , err: %s", err)
		os.Exit(1)
	}
}
