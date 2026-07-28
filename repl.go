package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CleanInput(text string) []string {
	var final []string

	if len(text) < 1 {
		return final
	}
	text = strings.ToLower(text)
	final = strings.Fields(text)
	return final

}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands map[string]cliCommand

func init() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Close the pokedex",
			callback:    commandExit,
		},

		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandhelp,
		},
	}
}

func Startrepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := CleanInput(scanner.Text())

		if len(input) == 0 {
			continue
		}

		if value, ok := commands[input[0]]; ok {
			value.callback()
		} else {
			fmt.Println("Unknown command\nUse 'help' for usage")
		}
	}
}

func commandExit() error {
	fmt.Println("Closing the pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandhelp() error {
	fmt.Print("Welcome to the Pokedex!\n\nUsage:\n")

	for key, value := range commands {
		fmt.Printf("\n%s: %s\n", key, value.description)
	}
	return nil
}
