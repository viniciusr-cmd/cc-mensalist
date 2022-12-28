package assettypes

import (
	"fmt"

	"github.com/goledgerdev/cc-tools/assets"
)

// Description of a Agency as a collection of vehicles
var Agency = assets.AssetType{
	Tag:         "agency",
	Label:       "Agency",
	Description: "Agency as a collection of vehicles",

	Props: []assets.AssetProp{
		{
			// Primary Key
			Required: true,
			IsKey:    true,
			Tag:      "cnpj",
			Label:    "CNPJ",
			DataType: "cnpj",
			Writers:  []string{`org3MSP`, "orgMSP"}, // This means only org3 can create the asset (others can edit)
		},
		{
			Required: true,
			Tag:      "name",
			Label:    "Agency rent name",
			DataType: "string",
		},
		{
			// Mandatory property
			Required: false,
			Tag:      "city",
			Label:    "City",
			DataType: "string",
			Validate: func(city interface{}) error {
				cityStr := city.(string)
				if cityStr == "" {
					return fmt.Errorf("city must be non-empty")
				}
				return nil
			},
		},
		{
			// Asset reference list
			Tag:      "vehicleList",
			Label:    "Vehicles List",
			DataType: "[]->vehicle",
		},
	},
}
