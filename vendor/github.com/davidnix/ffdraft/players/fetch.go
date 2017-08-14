package players

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const CacheLocation = "./cached_players.json"

// LoadPlayers gets data from http://fantasyfootballanalytics.net and also combines data from another source to
// update player injury and suspension notes
func Load() ([]Player, error) {
	resp, err := http.Post(analyticsURL, "application/json", strings.NewReader(analyticsRequestBody))
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("Bad response from server, got status code" + resp.Status)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var p Response
	if err := json.Unmarshal(data, &p); err != nil {
		return nil, err
	}

	if fileErr := ioutil.WriteFile(CacheLocation, data, 0644); fileErr != nil {
		fmt.Println("unable to cache response to disk")
	}

	players := p.Data.Players
	if len(players) == 0 {
		return nil, errors.New("response returned 0 players")
	}

	return players, nil
}

func LoadFromFile(path string) ([]Player, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var serialized Response
	if err := json.Unmarshal(data, &serialized); err != nil {
		return nil, err
	}
	return serialized.Data.Players, nil
}
