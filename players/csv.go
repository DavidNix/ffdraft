package players

import (
	"os"

	"bytes"
	"encoding/json"

	"math"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"github.com/pkg/errors"
)

// LoadFromCSV pre-processes a csv manually downloaded from http://apps.fantasyfootballanalytics.net/lineupoptimizer/.
// Login and use the download button to get the csv. Unfortunately, there lacks an easy way to make a request to get the data
// IMPORTANT: You want the custom rankings (not the raw).
// This function take the csv and transforms it into a parseable json file
func LoadFromCSV(path string) ([]Player, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	df := dataframe.ReadCSV(f,
		dataframe.DetectTypes(true),
		dataframe.WithTypes(map[string]series.Type{
			"cost": series.Float,
		}),
	)
	df.Capply(zeroInf)
	// fmt.Println(df)
	if df.Err != nil {
		return nil, errors.WithStack(df.Err)
	}
	buf := new(bytes.Buffer)
	if err := df.WriteJSON(buf); err != nil {
		return nil, errors.WithStack(err)
	}
	p := make([]Player, df.Nrow())
	if err := json.Unmarshal(buf.Bytes(), &p); err != nil {
		return nil, errors.WithStack(err)
	}
	return p, nil
}

func zeroInf(s series.Series) series.Series {
	for i := 0; i < s.Len(); i++ {
		val := s.Elem(i).Float()
		if math.IsInf(val, 0) {
			s.Set(i, series.New(0, s.Type(), s.Name))
		}
	}
	return s.Set(s.IsNaN(), series.New(0, s.Type(), s.Name))
}
