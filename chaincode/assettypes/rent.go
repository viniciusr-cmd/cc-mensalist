package assettypes

import (
	"github.com/goledgerdev/cc-mensalist/chaincode/datatypes"

	"github.com/goledgerdev/cc-tools/assets"
)

var Rent = assets.AssetType{
	Tag:         "rent",
	Label:       "Rent",
	Description: "Rent data",
	Props: []assets.AssetProp{
		{
			// Primary key
			Required: true,
			IsKey:    true,
			Tag:      "id",
			Label:    "Data ID",
			DataType: "string",
			Writers:  []string{`org3MSP`},
		},
		{
			// Property with default value reference another datatype
			Tag:          "status",
			Label:        "Rent status",
			DefaultValue: datatypes.NotRented,
			DataType:     "rentStatus",
		},
		{
			Tag:      "vehicle",
			Label:    "Vehicle for this rent",
			DataType: "->vehicle",
		},
		{
			Tag:      "renter",
			Label:    "Renter of this vehicle",
			DataType: "->renter",
		},
		{
			Tag:      "rentDate",
			Label:    "Rent date",
			DataType: "datetime",
		},
		{
			Tag:      "endDate",
			Label:    "End date",
			DataType: "datetime",
		},
	},
}
