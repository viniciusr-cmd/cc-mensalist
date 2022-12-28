package txdefs

import (
	"encoding/json"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	sw "github.com/goledgerdev/cc-tools/stubwrapper"
	tx "github.com/goledgerdev/cc-tools/transactions"
)

// POST Method
var CreateNewRenter = tx.Transaction{
	Tag:         "createNewRenter",
	Label:       "Create new renter",
	Description: "Create new renter",
	Method:      "POST",

	Args: []tx.Argument{
		{
			// Primary key
			Required: true,
			Tag:      "id",
			Label:    "CPF (Brazilian ID)",
			DataType: "cpf", // Datatypes are identified at datatypes folder
		},
		{
			// Mandatory property
			Required: true,
			Tag:      "renterName",
			Label:    "Renter name",
			DataType: "string",
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		cpf, _ := req["id"].(string)
		renterName, _ := req["renterName"].(string)

		renterMap := make(map[string]interface{})

		renterMap["@assetType"] = "renter"
		renterMap["id"] = cpf
		renterMap["renterName"] = renterName

		renterAsset, err := assets.NewAsset(renterMap)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to create a new renter asset")
		}

		_, err = renterAsset.PutNew(stub)
		if err != nil {
			return nil, errors.WrapError(err, "Error saving asset on blockchain")
		}

		renterJSON, nerr := json.Marshal(renterAsset)
		if nerr != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}

		return renterJSON, nil
	},
}
