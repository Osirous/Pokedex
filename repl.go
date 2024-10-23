package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startLoop() {
	input := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex >")

		input.Scan()
		text := input.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Unknown command. Try 'help' for valid commands.")
			continue
		}

		command.callback()
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints the help menu.",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Turns off the Pokedex.",
			callback:    callbackExit,
		},
		"map": {
			name:        "map",
			description: "Shows a list of areas.",
			callback:    callbackMap,
		},
	}
}

func cleanInput(str string) []string {
	lowerCase := strings.ToLower(str)
	words := strings.Fields(lowerCase)
	return words
}
