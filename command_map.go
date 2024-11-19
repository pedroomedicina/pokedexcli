package main

import (
	"errors"
	"fmt"
)

func callbackMap(cfg *config) error {
	resp, err := cfg.pokeApiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}

	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf("- %s\n", area.Name)
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.previousLocationAreaURL = resp.Previous
	return nil
}

func callbackMapb(cfg *config) error {
	if cfg.previousLocationAreaURL == nil {
		return errors.New("you are on the first page")
	}
	resp, err := cfg.pokeApiClient.ListLocationAreas(cfg.previousLocationAreaURL)
	if err != nil {
		return err
	}

	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf("- %s\n", area.Name)
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.previousLocationAreaURL = resp.Previous
	return nil
}
