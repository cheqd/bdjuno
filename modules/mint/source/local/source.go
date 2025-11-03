package local

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"cosmossdk.io/math"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/forbole/juno/v6/node/local"

	mintsource "github.com/forbole/callisto/v4/modules/mint/source"
)

var _ mintsource.Source = &Source{}

// Source implements mintsource.Source using a local node
type Source struct {
	*local.Source
	querier minttypes.QueryServer
}
type InflationResponse struct {
	Inflation string `json:"inflation"`
}

var out = &InflationResponse{}

// NewSource returns a new Source instance
func NewSource(source *local.Source, querier minttypes.QueryServer) *Source {
	return &Source{
		Source:  source,
		querier: querier,
	}
}

// GetInflation implements mintsource.Source
func (s Source) GetInflation(height int64) (math.LegacyDec, error) {
	restAddr := strings.TrimRight(os.Getenv("REST_ADDRESS"), "/")
	if restAddr == "" {
		return math.LegacyDec{}, fmt.Errorf("REST_ADDRESS not set")
	}

	url := restAddr + "/cosmos/mint/v1beta1/inflation"
	resp, err := http.Get(url)
	if err != nil {
		return math.LegacyDec{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return math.LegacyDec{}, fmt.Errorf("inflation REST query failed: %s - %s", resp.Status, string(body))
	}

	var out struct {
		Inflation string `json:"inflation"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return math.LegacyDec{}, err
	}
	if out.Inflation == "" {
		return math.LegacyDec{}, fmt.Errorf("empty inflation from REST")
	}

	dec, err := math.LegacyNewDecFromStr(out.Inflation)
	if err != nil {
		return math.LegacyDec{}, err
	}
	return dec, nil
}

// Params implements mintsource.Source
func (s Source) Params(height int64) (minttypes.Params, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return minttypes.Params{}, fmt.Errorf("error while loading height: %s", err)
	}

	res, err := s.querier.Params(ctx, &minttypes.QueryParamsRequest{})
	if err != nil {
		return minttypes.Params{}, err
	}

	return res.Params, nil
}
