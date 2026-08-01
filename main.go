package main

import (
	"time"
	"github.com/montruh-afk/pokedex/internal/pokeapi"
) 


func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, 1 * time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
		Pokedex: make(map[string]pokeapi.Pokemon),
	}

	Startrepl(cfg)
	
}
