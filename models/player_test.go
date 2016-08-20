package models

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	data, err := ioutil.ReadFile("../test/ff_response_fixture.json")
	if err != nil {
		assert.NoError(t, err)
		return
	}
	var serialized Response
	if err := json.Unmarshal(data, &serialized); err != nil {
		assert.NoError(t, err)
		return
	}
	players := serialized.Data.Players

	assert.True(t, len(players) > 0)

	donnie := players[0]
	assert.Equal(t, donnie.ID, 184)
	assert.Equal(t, donnie.Name, "Donnie Avery")
	assert.Equal(t, donnie.Position, "WR")
	assert.Equal(t, donnie.Team, "FA")
	assert.Equal(t, 10.0, donnie.TargetAuctionCost)
	assert.Equal(t, 9999.0, donnie.ECR)
	assert.Equal(t, 1543.0, donnie.OverallRank)
	assert.Equal(t, 226.0, donnie.PositionRank)
	assert.Equal(t, 2.0, donnie.Ceil)
	assert.Equal(t, 1.0, donnie.Floor)
	assert.Equal(t, 3.3124, donnie.Risk)
}
