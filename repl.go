package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startLoop(cfg *config) {
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
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Unknown command. Try 'help' for valid commands.")
			continue
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"catch": {
			name:        "catch {pokemon name}",
			description: "Your chance to capture a pokemon!",
			callback:    callbackCatch,
		},
		"pokedex": {
			name:        "pokedex",
			description: "A list of pokemon in your pokedex.",
			callback:    callbackPokedex,
		},
		"inspect": {
			name:        "inspect {pokemon name}",
			description: "Examine a captured pokemon!",
			callback:    callbackInspect,
		},
		"help": {
			name:        "help",
			description: "Prints the help menu.",
			callback:    callbackHelp,
		},
		"explore": {
			name:        "explore {location}",
			description: "Shows pokemon in area.",
			callback:    callbackExplore,
		},
		"exit": {
			name:        "exit",
			description: "Turns off the Pokedex.",
			callback:    callbackExit,
		},
		"map": {
			name:        "map",
			description: "Shows the next list of areas.",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Shows the previous list of areas.",
			callback:    callbackMapb,
		},
	}
}

func cleanInput(str string) []string {
	lowerCase := strings.ToLower(str)
	words := strings.Fields(lowerCase)
	return words
}
