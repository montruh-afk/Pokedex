package main

import (
	"fmt"
	"os"
	"errors"
	poke "github.com/montruh-afk/pokedex/internal/pokeapi"
 	random "math/rand"

)



type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

var commands map[string]cliCommand

func init() {
	
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Closes the Pokedex",
			callback:    commandExit,
		},

		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},

		"map": {
			name:        "map",
			description: "Each subsequent call to map will display the next 20 locations",
			callback:    commandMap,
		},

		"mapb": {
			name:        "mapb",
			description: "Returns the result of the most recent call to 'map'",
			callback:    commandMapb,
		},

		"explore": {
			name: "explore",
			description: "Lists all Pokémon located in selected area",
			callback: commandExplore,
		},

		"catch": {
			name: "catch",
			description: "Adds a Pokémon to the Pokedex",
			callback: commandCatch,
		},

		"inspect": {
			name: "inspect",
			description: "Returns information about a captured Pokémon",
			callback: commandInspect,
		},
	}
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Print("Welcome to the Pokedex!\n\nUsage:\n")

	for key, value := range commands {
		fmt.Printf("\n%s: %s\n", key, value.description)
	}
	return nil
}

func commandMap(cfg *config) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}
	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page\n")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandExplore(cfg *config) error {
	url := ""
	if cfg.prevLocationsURL == nil {
		url = fmt.Sprintf("%s%s/%s", poke.BaseURL, poke.DefaultLocation, *cfg.id)
	} else {
		url = fmt.Sprintf("%s/%s", *cfg.prevLocationsURL, *cfg.id)
	}
	locationsResp, err := cfg.pokeapiClient.ListLocations(&url)

	if err != nil {
		return err
	}
	

	fmt.Println("Exploring", *cfg.id)
	
	for _, pokemon := range locationsResp.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}

func commandCatch(cfg *config) error {
	if cfg.id == nil {
		return errors.New("you must provide a pokemon name")
	}
	id := *cfg.id
	
	pokemon, err := cfg.pokeapiClient.Catch(cfg.id)
	if err != nil {
		return fmt.Errorf("%s is not a valid Pokémon: %v", id, err)
	}
	fmt.Println("Throwing a Pokeball at", fmt.Sprintf("%s...", id))
	
	chance := random.Intn(300)
	if chance >= 299 {
		chance += 336
	}
	
	if chance > pokemon.BaseExperience {
		
		fmt.Println(id, "was caught!")
		cfg.Pokedex[id] = pokemon
		
	} else {
		fmt.Println(id, "escaped :(")
	}
	return nil
}

func commandInspect(cfg *config) error {
	if cfg.id == nil {
		return errors.New("you must provide a pokemon name to inspect\n")
	}
	id := *cfg.id
	if value, ok := cfg.Pokedex[id]; !ok {
		return fmt.Errorf("you have not caught %s\n", id)
	} else {
		poke.ParseStruct(&value, cfg.id)
	}
	return nil
}