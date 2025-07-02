package types

import (
	"github.com/cordialsys/go-sui-sdk/v2/lib"
	"github.com/cordialsys/go-sui-sdk/v2/sui_types"
)

type DynamicFieldInfo struct {
	Name sui_types.DynamicFieldName `json:"name"`
	//Base58
	BcsName    lib.Base58                              `json:"bcsName"`
	Type       lib.TagJson[sui_types.DynamicFieldType] `json:"type"`
	ObjectType string                                  `json:"objectType"`
	ObjectId   sui_types.ObjectID                      `json:"objectId"`
	Version    sui_types.SequenceNumber                `json:"version"`
	Digest     sui_types.ObjectDigest                  `json:"digest"`
}

type DynamicFieldPage = Page[DynamicFieldInfo, string]
