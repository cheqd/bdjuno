package remote

import (
	"context"
	"io"

	"github.com/cosmos/cosmos-sdk/codec"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	gogoproto "github.com/cosmos/gogoproto/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding"
)

func init() {
	// Register gogoproto codec for gRPC
	encoding.RegisterCodec(gogoprotoCodec{})
}

// gogoprotoCodec implements grpc/encoding.Codec using gogoproto
type gogoprotoCodec struct{}

func (gogoprotoCodec) Name() string {
	return "gogoproto"
}

func (gogoprotoCodec) Marshal(v interface{}) ([]byte, error) {
	if msg, ok := v.(gogoproto.Marshaler); ok {
		return msg.Marshal()
	}
	return nil, io.EOF
}

func (gogoprotoCodec) Unmarshal(data []byte, v interface{}) error {
	if msg, ok := v.(gogoproto.Unmarshaler); ok {
		return msg.Unmarshal(data)
	}
	return io.EOF
}

// grpcConnWrapper wraps a standard grpc.ClientConn to implement grpc1.ClientConn interface
// This allows us to use gogoproto-generated clients with standard gRPC connections
type grpcConnWrapper struct {
	*grpc.ClientConn
}

// NewStream implements grpc1.ClientConn interface
func (w *grpcConnWrapper) NewStream(ctx context.Context, desc *grpc.StreamDesc, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	return w.ClientConn.NewStream(ctx, desc, method, opts...)
}

// Invoke implements grpc1.ClientConn interface
// Use gogoproto codec via CallOption
func (w *grpcConnWrapper) Invoke(ctx context.Context, method string, args, reply interface{}, opts ...grpc.CallOption) error {
	opts = append(opts, grpc.ForceCodec(gogoprotoCodec{}))
	return w.ClientConn.Invoke(ctx, method, args, reply, opts...)
}

func (w *grpcConnWrapper) Close() error {
	return w.ClientConn.Close()
}

// GogoprotoQueryClient uses gogoproto-generated types for mint queries
type GogoprotoQueryClient struct {
	cc  *grpc.ClientConn
	cdc codec.Codec
	pb  minttypes.QueryClient
}

// NewGogoprotoQueryClient creates a new gogoproto-based query client
func NewGogoprotoQueryClient(cc *grpc.ClientConn, cdc codec.Codec) *GogoprotoQueryClient {
	// Wrap the standard grpc.ClientConn to implement grpc1.ClientConn interface
	wrapped := &grpcConnWrapper{ClientConn: cc}

	// Create the gogoproto-generated client which uses gogoproto for marshaling/unmarshaling
	pbClient := minttypes.NewQueryClient(wrapped)

	return &GogoprotoQueryClient{
		cc:  cc,
		cdc: cdc,
		pb:  pbClient,
	}
}

func (c *GogoprotoQueryClient) Inflation(ctx context.Context, req *minttypes.QueryInflationRequest, opts ...grpc.CallOption) (*minttypes.QueryInflationResponse, error) {
	return c.pb.Inflation(ctx, req, opts...)
}

func (c *GogoprotoQueryClient) Params(ctx context.Context, req *minttypes.QueryParamsRequest, opts ...grpc.CallOption) (*minttypes.QueryParamsResponse, error) {
	return c.pb.Params(ctx, req, opts...)
}
