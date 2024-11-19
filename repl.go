package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex > ")
		if !scanner.Scan() {
			fmt.Println("Exiting...")
			break
		}

		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]
		availableCommands := getCommands()
		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Printf("Unknown command: %s\n", commandName)
			continue
		}

		command.callback(cfg)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    callbackExit,
		},
		"map": {
			name:        "map",
			description: "Lists some location areas, consecutive calls will paginate forward",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Go to a previous location area page",
			callback:    callbackMapb,
		},
	}
}
