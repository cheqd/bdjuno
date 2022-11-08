package handlers

import (
	"fmt"
	"strings"

	"github.com/forbole/bdjuno/v3/modules/actions/types"
	"github.com/rs/zerolog/log"
)

func CheqdDIDAnalytics(ctx *types.Context, payload *types.Payload) (interface{}, error) {
	opts := payload.GetAnalyticsOptions()

	log.Debug().Str("action", "CheqdDIDAnalytics").
		Str("msgTypes", opts.MsgTypes).
		Send()

	msgTypes := strings.Split(opts.MsgTypes, ",")
	if invalidMsgType, ok := validateDIDMsgTypes(msgTypes); !ok {
		return nil, fmt.Errorf("error: invalid message type - %s", invalidMsgType)
	}

	resp, err := ctx.DB.GetTotalDIDs(msgTypes, opts.Plottable)
	if err != nil {
		return nil, fmt.Errorf("error while getting DID analytics: %w", err)
	}

	return resp, nil
}

func CheqdResourceAnalytics(ctx *types.Context, payload *types.Payload) (interface{}, error) {
	opts := payload.GetAnalyticsOptions()

	log.Debug().Str("action", "CheqdResourceAnalytics").
		Str("msgTypes", opts.MsgTypes).
		Send()

	msgTypes := strings.Split(opts.MsgTypes, ",")
	if invalidMsgType, ok := validateResourceMsgTypes(msgTypes); !ok {
		return nil, fmt.Errorf("error: invalid message type - %s", invalidMsgType)
	}

	resp, err := ctx.DB.GetTotalResources(msgTypes, opts.Plottable)
	if err != nil {
		return nil, fmt.Errorf("error while getting resource analytics: %w", err)
	}

	return resp, nil
}

func validateDIDMsgTypes(msgTypes []string) (string, bool) {
	// Supported MsgTypes "MsgCreateDID" "MsgUpdateDID"

	if len(msgTypes) == 0 {
		return "", false
	}

	for _, msgType := range msgTypes {
		switch msgType {
		case "MsgCreateDid":
		case "MsgUpdateDid":
		default:
			return msgType, false
		}
	}

	return "", true
}

func validateResourceMsgTypes(msgTypes []string) (string, bool) {
	// Supported MsgTypes "MsgCreateResource"

	for _, msgType := range msgTypes {
		switch msgType {
		case "MsgCreateResource":
		default:
			return msgType, false
		}
	}

	return "", true
}
