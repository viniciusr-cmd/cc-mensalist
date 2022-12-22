package main

import (
	"github.com/goledgerdev/cc-mensalist/chaincode/assettypes"
	"github.com/goledgerdev/cc-tools/assets"
)

var assetTypeList = []assets.AssetType{
	assettypes.Vehicle,
	assettypes.Agency,
	assettypes.Renter,
	assettypes.Rent,
}
