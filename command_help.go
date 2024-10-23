package main

import "fmt"

func callbackHelp(cfg *config, args ...string) error {
	fmt.Println("")
	fmt.Println("These are the currently available commands:")

	availableCommands := getCommands()
	for _, cmd := range availableCommands {
		fmt.Printf(" * %s: %s\n", cmd.name, cmd.description)
	}

	fmt.Println("")

	return nil
}
