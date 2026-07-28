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
			continue
		}

		if value, ok := commands[input[0]]; ok {
			err := value.callback(cfg)
			if err != nil {
				fmt.Print(err)
			}
		} else {
			fmt.Println("Unknown command - Use 'help' for usage")
		}
	}
}