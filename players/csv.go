package players

import (
	"encoding/json"
	"io/ioutil"
)

const cacheLocation = "./cached_players.json"

// LoadFromCSV pre-processes a csv manually downloaded from http://apps.fantasyfootballanalytics.net/lineupoptimizer/.
// Login and use the download button to get the csv. Unfortunately, there lacks an easy way to make a request to get the data
// IMPORTANT: You want the custom rankings (not the raw).
// This function take the csv and transforms it into a parseable json file
func LoadFromCSV(path string) ([]Player, error) {
	data, err := ioutil.ReadFile(cacheLocation)
	if err != nil {
		return nil, err
	}
	var serialized Response
	if err := json.Unmarshal(data, &serialized); err != nil {
		return nil, err
	}
	return serialized.Data.Players, nil
}
