package remote

import (
	"context"

	"cosmossdk.io/math"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	mintsource "github.com/forbole/callisto/v4/modules/mint/source"
	"github.com/forbole/juno/v6/node/remote"
	"google.golang.org/grpc"
)

var _ mintsource.Source = &Source{}

// Source implements mintsource.Source using a remote node with gogoproto-compatible types
type Source struct {
	*remote.Source
	querier QueryClient
}

// QueryClient interface for mint queries using gogoproto-generated types
type QueryClient interface {
	Inflation(ctx context.Context, req *minttypes.QueryInflationRequest, opts ...grpc.CallOption) (*minttypes.QueryInflationResponse, error)
	Params(ctx context.Context, req *minttypes.QueryParamsRequest, opts ...grpc.CallOption) (*minttypes.QueryParamsResponse, error)
}

// NewSource returns a new Source instance
func NewSource(source *remote.Source, querier QueryClient) *Source {
	return &Source{
		Source:  source,
		querier: querier,
	}
}

// GetInflation implements mintsource.Source
func (s Source) GetInflation(height int64) (math.LegacyDec, error) {
	req := &minttypes.QueryInflationRequest{}
	res, err := s.querier.Inflation(remote.GetHeightRequestContext(s.Ctx, height), req)
	if err != nil {
		return math.LegacyDec{}, err
	}

	return res.Inflation, nil
}

// Params implements mintsource.Source
func (s Source) Params(height int64) (minttypes.Params, error) {
	req := &minttypes.QueryParamsRequest{}
	res, err := s.querier.Params(remote.GetHeightRequestContext(s.Ctx, height), req)
	if err != nil {
		return minttypes.Params{}, err
	}

	// res.Params is already of type minttypes.Params (proto-generated types are canonical)
	return res.Params, nil
}
