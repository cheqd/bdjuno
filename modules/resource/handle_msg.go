package resource

import (
	"fmt"

	resourcetypes "github.com/cheqd/cheqd-node/x/resource/types"
	"github.com/forbole/callisto/v4/types"
	"github.com/forbole/callisto/v4/utils"
	juno "github.com/forbole/juno/v6/types"
	"github.com/rs/zerolog/log"
)

var msgFilter = map[string]bool{
	"/cheqd.resource.v2.MsgCreateResource": true,
}

// HandleMsg implements MessageModule
func (m *Module) HandleMsg(index int, msg juno.Message, tx *juno.Transaction) error {
	if _, ok := msgFilter[msg.GetType()]; !ok {
		return nil
	}

	log.Debug().Str("module", "resource").Str("hash", tx.TxHash).Uint64("height", tx.Height).Msg(fmt.Sprintf("handling resource message %s", msg.GetType()))

	switch msg.GetType() {
	case "/cheqd.resource.v2.MsgCreateResource":
		cosmosMsg := utils.UnpackMessage(m.cdc, msg.GetBytes(), &resourcetypes.MsgCreateResource{})
		return m.handleMsgCreateResource(int64(tx.Height), cosmosMsg, tx.FeePayer(m.cdc))
	default:
		return nil
	}
}

func (m *Module) handleMsgCreateResource(height int64, msg *resourcetypes.MsgCreateResource, feePayer []byte) error {
	feePayerAddr, err := m.cdc.InterfaceRegistry().SigningContext().AddressCodec().BytesToString(feePayer)
	if err != nil {
		return err
	}

	return m.db.SaveResource(types.NewResource(msg.Payload.Id, msg.Payload.CollectionId,
		msg.Payload.Data, msg.Payload.Name, msg.Payload.Version,
		msg.Payload.ResourceType, msg.Payload.AlsoKnownAs, feePayerAddr, height))
}
