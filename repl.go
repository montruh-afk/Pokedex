package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"encoding/json"
	"github.com/montruh-afk/pokedex/internal/pokeapi"
	"path/filepath"
	"errors"
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
	Load(cfg)

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

func Save(cfg *config) error {
	saveFile, err := json.Marshal(cfg.Pokedex)
	if err != nil {
		return err
	}

	filePath, err := getWorkingDir()
	if err != nil {
		return fmt.Errorf("error getting working directory: %v", err)
	}

	fullPath := filepath.Join(filePath, "data.json")

	// Use os.OpenFile to either create or open the file for writing
	file, err := os.OpenFile(fullPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	_, err = file.Write(saveFile)
	if err != nil {
		return fmt.Errorf("something went wrong writing data: %v", err)
	}

	return nil
}

func Load(cfg *config) error {
	filePath, err := getWorkingDir()
	if err != nil {
		return fmt.Errorf("error getting working directory: %v", err)
	}

	fullPath := filepath.Join(filePath, "data.json")

	// Check if the file exists using os.Stat and errors.Is
	_, err = os.Stat(fullPath)
	if errors.Is(err, os.ErrNotExist) {
		// File doesn't exist yet; initialize an empty map and return no error
		if cfg.Pokedex == nil {
			cfg.Pokedex = make(map[string]pokeapi.Pokemon)
		}
		return nil 
	} else if err != nil {
		return fmt.Errorf("error checking save file: %v", err)
	}

	// Read the entire file using os.ReadFile
	data, err := os.ReadFile(fullPath)
	if err != nil {
		return fmt.Errorf("error reading save file: %v", err)
	}

	// Ensure the map is initialized before unmarshalling
	if cfg.Pokedex == nil {
		cfg.Pokedex = make(map[string]pokeapi.Pokemon)
	}

	// Parse the JSON directly back into the Pokedex map
	if err := json.Unmarshal(data, &cfg.Pokedex); err != nil {
		return fmt.Errorf("error parsing save data: %v", err)
	}

	return nil
}
