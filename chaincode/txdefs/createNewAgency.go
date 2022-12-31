package txdefs

import (
	"encoding/json"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	sw "github.com/goledgerdev/cc-tools/stubwrapper"
	tx "github.com/goledgerdev/cc-tools/transactions"
)

// Create a new Agency on channel
// POST Method
var CreateNewAgency = tx.Transaction{
	Tag:         "createNewAgency",
	Label:       "Create New Agency",
	Description: "Create a New Agency",
	Method:      "POST",
	Callers:     []string{"$org3MSP", "$orgMSP"}, // Only org3 can call this transaction

	Args: []tx.Argument{
		{
			Tag:         "name",
			Label:       "Name",
			Description: "Name of the Agency",
			DataType:    "string",
			Required:    true,
		},
		{
			Tag:         "cnpj",
			Label:       "CNPJ",
			Description: "Agency's CNPJ",
			DataType:    "cnpj",
			Required:    true,
		},
		{
			Tag:         "city",
			Label:       "city",
			Description: "Agency's City",
			DataType:    "string",
			Required:    true,
		},
		{
			Tag:         "vehicleList",
			Label:       "Vehicles For Rent",
			Description: "vehicles",
			DataType:    "[]->vehicle",
			Required:    false,
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		name, _ := req["name"].(string)
		cnpj, _ := req["cnpj"].(string)
		city, _ := req["city"].(string)
		vehicles, _ := req["vehicleList"].([]interface{})

		agencyMap := make(map[string]interface{})

		agencyMap["@assetType"] = "agency"
		agencyMap["cnpj"] = cnpj
		agencyMap["name"] = name
		agencyMap["city"] = city

		if len(vehicles) > 0 {
			agencyMap["vehicleList"] = vehicles
		}

		agencyAsset, err := assets.NewAsset(agencyMap)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to create a new asset")
		}

		// Save the new agency on channel
		_, err = agencyAsset.PutNew(stub)
		if err != nil {
			return nil, errors.WrapError(err, "Error saving asset on blockchain")
		}

		// Marshal asset back to JSON format
		agencyJSON, nerr := json.Marshal(agencyAsset)
		if nerr != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}

		return agencyJSON, nil
	},
}
