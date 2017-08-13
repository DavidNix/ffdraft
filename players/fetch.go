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

// PreprocessCSV pre-processes a csv manually downloaded from http://apps.fantasyfootballanalytics.net/lineupoptimizer/.
// Login and use the download button to get the csv. Unfortunately, there lacks an easy way to make a request to get the data
// IMPORTANT: You want the custom rankings (not the raw).
// This function take the csv and transforms it into a parseable json file
func PreprocessCSV(path string) error {

	return nil
}

func LoadFromCSV(path string) ([]Player, error) {
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
