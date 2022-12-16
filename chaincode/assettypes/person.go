package assettypes

import (
	"fmt"

	"github.com/goledgerdev/cc-tools/assets"
)

var Mensalist = assets.AssetType{
	Tag:         "mensalist",
	Label:       "Mensalist",
	Description: "Personal data of mensalist",

	Props: []assets.AssetProp{
		{
			// Primary key
			Required: true,
			IsKey:    true,
			Tag:      "id",
			Label:    "CPF (Brazilian ID)",
			DataType: "cpf",                         // Datatypes are identified at datatypes folder
			Writers:  []string{`org1MSP`, "orgMSP"}, // This means only org1 can create the asset (others can edit)
		},
		{
			// Mandatory property
			Required: false,
			Tag:      "CNPJ",
			Label:    "CNPJ (Brazilian National Registry of Legal Entities)",
			DataType: "cnpj",
		},
		{
			// Mandatory property
			Required: true,
			Tag:      "name",
			Label:    "Mensalist name",
			DataType: "string",
			// Validate funcion
			Validate: func(name interface{}) error {
				nameStr := name.(string)
				if nameStr == "" {
					return fmt.Errorf("name must be non-empty")
				}
				return nil
			},
		},
		{
			Required: true,
			Tag:      "Registry date",
			Label:    "Mensalist registry date",
			DataType: "datetime",
		},
		{
			Required: true,
			Tag:      "Vehicles",
			Label:    "Vehicles of mensalist",
			DataType: "[]->vehicle",
		},
	},
}
