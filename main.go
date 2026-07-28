package main

import (
	"time"
	"github.com/montruh-afk/pokedex/internal/pokeapi"
) 


func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	Startrepl(cfg)
	
}
