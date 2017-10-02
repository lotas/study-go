package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type User struct {
	Name string
	ID   int
}

type Nail struct {
	rating       int
	countAnswers int
	countVotes   int
	Text         string
	id           int
	created      time.Time
	User         *User `json:"$user"`
}

type Joke struct {
	ID         int    `json:"id"`
	Text       string `json:"text"`
	countVotes int
	rating     int
	created    time.Time
	nsfw       int
	Nail       *Nail `json:"$nail"`
	User       *User `json:"$user"`
}

type ResponsePager struct {
	Pages int
	Total int
	Page  int
	First string
	Last  string
	Next  string
	Prev  string
}

type TopJokesResponse struct {
	Jokes []Joke
	Pager ResponsePager
}

type Status struct {
	Started string  `json:"started"`
	Uptime  float32 `json:"uptime"`
}

func main() {

	// st, _ := fetchStatus()
	// fmt.Printf("Status: started: %+v, uptime: %+v\n", st.Started, st.Uptime)

	// fmt.Println("Fetching TOP jokes\n")

	data, err := fetchTopJokes()
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("%+v", data)
	// fmt.Printf("Pager: %+v", data.Pager)

	for _, joke := range data.Jokes[:4] {
		fmt.Printf("#%d: (%s, %s)\n- %s\n- %s\n\n",
			joke.ID, joke.Nail.User.Name, joke.User.Name,
			joke.Nail.Text, joke.Text)
	}
}

const apiUrl = "https://api.loliful.io"

func fetchStatus() (*Status, error) {
	resp, err := http.Get(apiUrl + "/status")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Cannot fetch status: %s", resp.Status)
	}

	var status Status
	// var info map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&status); err != nil {
		return nil, err
	}

	// fmt.Printf("%+v\n\n", info)

	return &status, nil
}

func fetchTopJokes() (*TopJokesResponse, error) {
	resp, err := http.Get(apiUrl + "/top")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Cannot fetch jokes: %s", resp.Status)
	}

	var jokeResp TopJokesResponse
	// var info map[string]interface{}

	if err := json.NewDecoder(resp.Body).Decode(&jokeResp); err != nil {
		return nil, err
	}
	// fmt.Printf("%+v", info)

	return &jokeResp, nil
}
