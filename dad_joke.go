package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// DadJoke represents a dad joke from the API.
type DadJoke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

// DadJokeHandler handles the command of getting a joke.
func DadJokeHandler() error {
	req, err := http.NewRequest("GET", "https://icanhazdadjoke.com", nil)
	if err != nil {
		return fmt.Errorf("creating request: %v", err)
	}
	req.Header.Add("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("sending request: %v", err)
	}
	defer resp.Body.Close()

	var dadJoke *DadJoke
	if err := json.NewDecoder(resp.Body).Decode(&dadJoke); err != nil {
		return fmt.Errorf("parsing request body: %v", err)
	}

	fmt.Fprintln(os.Stdout, dadJoke.Joke)

	return nil
}
