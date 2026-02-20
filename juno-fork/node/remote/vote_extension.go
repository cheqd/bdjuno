package remote

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
)

// InjectedVoteExtensionTx matches cheqd.oracle.v2.InjectedVoteExtensionTx
// This is the vote extension tx injected by the prepare proposal handler
type InjectedVoteExtensionTx struct {
	ExchangeRateVotes  []AggregateExchangeRateVote `protobuf:"bytes,1,rep,name=exchange_rate_votes,json=exchangeRateVotes,proto3" json:"exchange_rate_votes"`
	ExtendedCommitInfo []byte                      `protobuf:"bytes,2,opt,name=extended_commit_info,json=extendedCommitInfo,proto3" json:"extended_commit_info,omitempty"`
}

func (m *InjectedVoteExtensionTx) Reset()         { *m = InjectedVoteExtensionTx{} }
func (m *InjectedVoteExtensionTx) String() string { return proto.CompactTextString(m) }
func (*InjectedVoteExtensionTx) ProtoMessage()    {}

// AggregateExchangeRateVote matches cheqd.oracle.v2.AggregateExchangeRateVote
type AggregateExchangeRateVote struct {
	ExchangeRates sdk.DecCoins `protobuf:"bytes,1,rep,name=exchange_rates,json=exchangeRates,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"exchange_rates"`
	Voter         string       `protobuf:"bytes,2,opt,name=voter,proto3" json:"voter,omitempty"`
}

func (m *AggregateExchangeRateVote) Reset()         { *m = AggregateExchangeRateVote{} }
func (m *AggregateExchangeRateVote) String() string { return proto.CompactTextString(m) }
func (*AggregateExchangeRateVote) ProtoMessage()    {}

// DecodeVoteExtension attempts to decode raw vote extension bytes using protobuf
func DecodeVoteExtension(data []byte) (*InjectedVoteExtensionTx, error) {
	if len(data) < 10 {
		return nil, fmt.Errorf("data too short: %d bytes", len(data))
	}

	var voteExt InjectedVoteExtensionTx
	if err := proto.Unmarshal(data, &voteExt); err != nil {
		return nil, fmt.Errorf("failed to unmarshal vote extension: %w", err)
	}

	if len(voteExt.ExchangeRateVotes) == 0 {
		return nil, fmt.Errorf("no exchange rate votes found")
	}

	return &voteExt, nil
}
