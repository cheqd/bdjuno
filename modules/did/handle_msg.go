package did

import (
	"fmt"

	didtypes "github.com/cheqd/cheqd-node/x/did/types"
	"github.com/rs/zerolog/log"

	"github.com/forbole/callisto/v4/types"
	"github.com/forbole/callisto/v4/utils"
	juno "github.com/forbole/juno/v6/types"
)

var msgFilter = map[string]bool{
	"/cheqd.did.v2.MsgCreateDidDoc":     true,
	"/cheqd.did.v2.MsgUpdateDidDoc":     true,
	"/cheqd.did.v2.MsgDeactivateDidDoc": true,
}

// HandleMsg implements MessageModule
func (m *Module) HandleMsg(index int, msg juno.Message, tx *juno.Transaction) error {
	if _, ok := msgFilter[msg.GetType()]; !ok {
		return nil
	}

	log.Debug().Str("module", "did").Str("hash", tx.TxHash).Uint64("height", tx.Height).Msg(fmt.Sprintf("handling did message %s", msg.GetType()))

	switch msg.GetType() {
	case "/cheqd.did.v2.MsgCreateDidDoc":
		cosmosMsg := utils.UnpackMessage(m.cdc, msg.GetBytes(), &didtypes.MsgCreateDidDoc{})
		return m.handleMsgCreateDidDoc(int64(tx.Height), cosmosMsg, tx.FeePayer(m.cdc))

	case "/cheqd.did.v2.MsgUpdateDidDoc":
		cosmosMsg := utils.UnpackMessage(m.cdc, msg.GetBytes(), &didtypes.MsgUpdateDidDoc{})
		return m.handleMsgUpdateDidDoc(int64(tx.Height), cosmosMsg, tx.FeePayer(m.cdc))

	case "/cheqd.did.v2.MsgDeactivateDidDoc":
		cosmosMsg := utils.UnpackMessage(m.cdc, msg.GetBytes(), &didtypes.MsgDeactivateDidDoc{})
		return m.handleMsgDeactivateDidDoc(cosmosMsg)
	}

	return nil
}

func (m *Module) handleMsgCreateDidDoc(height int64, msg *didtypes.MsgCreateDidDoc, feePayer []byte) error {
	feePayerAddr, err := m.cdc.InterfaceRegistry().SigningContext().AddressCodec().BytesToString(feePayer)
	if err != nil {
		return err
	}

	return m.db.SaveDidDoc(types.NewDidDoc(msg.Payload.Id, msg.Payload.Context,
		msg.Payload.Controller, msg.Payload.VerificationMethod, msg.Payload.Authentication,
		msg.Payload.AssertionMethod, msg.Payload.CapabilityInvocation,
		msg.Payload.CapabilityDelegation, msg.Payload.KeyAgreement,
		msg.Payload.Service, msg.Payload.AlsoKnownAs, msg.Payload.VersionId, feePayerAddr, height))
}

func (m *Module) handleMsgUpdateDidDoc(height int64, msg *didtypes.MsgUpdateDidDoc, feePayer []byte) error {
	feePayerAddr, err := m.cdc.InterfaceRegistry().SigningContext().AddressCodec().BytesToString(feePayer)
	if err != nil {
		return err
	}

	return m.db.SaveDidDoc(types.NewDidDoc(msg.Payload.Id, msg.Payload.Context,
		msg.Payload.Controller, msg.Payload.VerificationMethod, msg.Payload.Authentication,
		msg.Payload.AssertionMethod, msg.Payload.CapabilityInvocation,
		msg.Payload.CapabilityDelegation, msg.Payload.KeyAgreement,
		msg.Payload.Service, msg.Payload.AlsoKnownAs, msg.Payload.VersionId, feePayerAddr, height))
}

func (m *Module) handleMsgDeactivateDidDoc(msg *didtypes.MsgDeactivateDidDoc) error {
	return m.db.DeleteDidDoc(msg.Payload.Id, msg.Payload.VersionId)
}
