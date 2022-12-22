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
			Tag:      "model",
			Label:    "Vehicle model name",
			DataType: "string",
		},
		{
			Tag:      "rentPrice",
			Label:    "Rent Price (BRL)",
			DataType: "number",
		},
		{
			Tag:      "renter",
			Label:    "Renter",
			DataType: "->renter",
		},
		{
			Tag:      "rent",
			Label:    "Rent ID",
			DataType: "->rent",
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
