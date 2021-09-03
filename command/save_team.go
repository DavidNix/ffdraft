package command

import (
	"encoding/json"
	"os"

	"github.com/davidnix/ffdraft/players"
	"github.com/fatih/color"
)

func SaveTeam(repo *players.Repo) {
	b, err := json.Marshal(repo.MyTeam)
	if err != nil {
		panic(err)
	}
	fname := repo.MyTeam.Name + ".json"
	err = os.WriteFile(fname, b, 0660)
	if err != nil {
		color.Red("Unable to save team file %q: %v", fname, err)
		return
	}
	color.Green("Your team was saved to %s", fname)
}
