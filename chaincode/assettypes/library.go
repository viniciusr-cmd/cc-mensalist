package assettypes

import (
	"fmt"

	"github.com/goledgerdev/cc-tools/assets"
)

// Description of a Library as a collection of books
var Agency = assets.AssetType{
	Tag:         "Agency",
	Label:       "Agency",
	Description: "Agency as a collection of vehicles",

	Props: []assets.AssetProp{
		{
			// Primary Key
			Required: true,
			IsKey:    true,
			Tag:      "name",
			Label:    "Agency rent name",
			DataType: "string",
			Writers:  []string{`org3MSP`, "orgMSP"}, // This means only org3 can create the asset (others can edit)
		},
		{
			Required: true,
			Tag:      "CNPJ",
			Label:    "CNPJ",
			DataType: "cnpj",
		},
		{
			// Mandatory property
			Required: true,
			Tag:      "city",
			Label:    "City",
			DataType: "string",
			Validate: func(city interface{}) error {
				cityStr := city.(string)
				if cityStr == "" {
					return fmt.Errorf("City must be non-empty")
				}
				return nil
			},
		},
		{
			// Asset reference list
			Tag:      "vehicles",
			Label:    "Vehicles for Rent",
			DataType: "[]->vehicle",
		},
		{
			// Asset reference list
			Tag:      "mensalist",
			Label:    "Mensalists registered",
			DataType: "->mensalist",
		},
	},
}
