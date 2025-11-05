package remote

import (
	"context"

	"cosmossdk.io/math"
	mintpb "github.com/cheqd/bdjuno/proto/cosmos/mint/v1beta1"
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
	Inflation(ctx context.Context, req *mintpb.QueryInflationRequest, opts ...grpc.CallOption) (*mintpb.QueryInflationResponse, error)
	Params(ctx context.Context, req *mintpb.QueryParamsRequest, opts ...grpc.CallOption) (*mintpb.QueryParamsResponse, error)
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
	req := &mintpb.QueryInflationRequest{}
	res, err := s.querier.Inflation(remote.GetHeightRequestContext(s.Ctx, height), req)
	if err != nil {
		return math.LegacyDec{}, err
	}

	return res.Inflation, nil
}

// Params implements mintsource.Source
func (s Source) Params(height int64) (minttypes.Params, error) {
	req := &mintpb.QueryParamsRequest{}
	res, err := s.querier.Params(remote.GetHeightRequestContext(s.Ctx, height), req)
	if err != nil {
		return minttypes.Params{}, err
	}

	// Convert from generated Params to cosmos-sdk minttypes.Params
	return convertParams(res.Params), nil
}

// convertParams converts from generated Params to cosmos-sdk minttypes.Params
func convertParams(p mintpb.Params) minttypes.Params {
	return minttypes.Params{
		MintDenom:           p.MintDenom,
		InflationRateChange: p.InflationRateChange,
		InflationMax:        p.InflationMax,
		InflationMin:        p.InflationMin,
		GoalBonded:          p.GoalBonded,
		BlocksPerYear:       p.BlocksPerYear,
	}
}
