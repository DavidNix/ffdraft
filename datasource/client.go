package datasource

import (
	"encoding/json"
	"errors"
	"github.com/davidnix/ffdraft/models"
	"io/ioutil"
	"net/http"
	"strings"
)

// LoadPlayers gets data from http://fantasyfootballanalytics.net and also combines data from another source to
// update player injury and suspension notes
func LoadPlayers() ([]models.Player, error) {
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
	var p models.Response
	if err := json.Unmarshal(data, &p); err != nil {
		return nil, err
	}

	return p.Data.Players, nil
}
