package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/montruh-afk/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	id *string
	Pokedex map[string]pokeapi.Pokemon
}

func CleanInput(text string) []string {
	var final []string

	if len(text) < 1 {
		return final
	}
	text = strings.ToLower(text)
	final = strings.Fields(text)
	return final

}
func Startrepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := CleanInput(scanner.Text())

		if len(input) == 0 {
			fmt.Println("Pokedex - see 'help' for usage")
		}

		if value, ok := commands[input[0]]; ok {
			if len(input) > 1 {
				cfg.id = &input[1]
			}
			err := value.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command - see 'help' for usage")
		}
	}
}
