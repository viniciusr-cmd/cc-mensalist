package assettypes

import (
	"github.com/goledgerdev/cc-tools/assets"
)

var Vehicle = assets.AssetType{
	Tag:         "vehicle",
	Label:       "Vehicle",
	Description: "Data of a vehicle",

	Props: []assets.AssetProp{
		{
			// Primary key
			Required: true,
			IsKey:    true,
			Tag:      "id",
			Label:    "Mercosur Plate of Vehicle",
			DataType: "plate",
			Writers:  []string{`org2MSP`}, // This means only org2 can create the asset (others can edit)
		},
		{
			//Mandatory property
			Required: true,
			Tag:      "type",
			Label:    "Vehicle type",
			DataType: "vehicleType",
		},
		{
			//Mandatory property
			Required: true,
			Tag:      "model",
			Label:    "Vehicle model",
			DataType: "string",
		},
		{
			Tag:      "mensalist",
			Label:    "Current mensalist",
			DataType: "->mensalist",
		},
		{
			//Mandatory property
			Required: true,
			Tag:      "rentPrice",
			Label:    "Rent price",
			DataType: "float",
		},
		{
			// Property with default value
			Tag:          "available",
			Label:        "Available for rent?",
			DefaultValue: true,
			DataType:     "boolean",
		},
	},
}
