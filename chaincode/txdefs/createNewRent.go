package txdefs

import (
	"encoding/json"

	"github.com/goledgerdev/cc-mensalist/chaincode/datatypes"
	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	sw "github.com/goledgerdev/cc-tools/stubwrapper"
	tx "github.com/goledgerdev/cc-tools/transactions"
)

// Create a new Agency on channel
// POST Method
var CreateNewRent = tx.Transaction{
	Tag:         "createNewRent",
	Label:       "Create New Rent",
	Description: "Create a New Rent",
	Method:      "POST",
	Callers:     []string{"$org2MSP", "$orgMSP"}, // Only org3 can call this transaction

	Args: []tx.Argument{
		{
			Tag:         "id",
			Label:       "Data ID",
			Description: "Rent ID",
			DataType:    "number",
			Required:    true,
		},
		{
			Tag:         "status",
			Label:       "Rent Status",
			Description: "Rent Status",
			DataType:    "rentStatus",
			Required:    true,
		},
		{
			Tag:         "vehicle",
			Label:       "Vehicle for this rent",
			Description: "Vehicle for rent",
			DataType:    "->vehicle",
			Required:    true,
		},
		{
			Tag:         "renter",
			Label:       "Renter",
			Description: "renter",
			DataType:    "->renter",
			Required:    true,
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		rentID, _ := req["id"].(int)
		rentStatus, _ := req["status"].(datatypes.RentStatus)
		vehicle, _ := req["vehicle"].(assets.Key)
		renter, _ := req["renter"].(assets.Key)

		rentMap := make(map[string]interface{})

		rentMap["@assetType"] = "rent"
		rentMap["id"] = rentID
		rentMap["status"] = rentStatus
		rentMap["vehicle"] = vehicle
		rentMap["renter"] = renter

		rentAsset, err := assets.NewAsset(rentMap)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to create a new asset")
		}

		// Save the new agency on channel
		_, err = rentAsset.PutNew(stub)
		if err != nil {
			return nil, errors.WrapError(err, "Error saving asset on blockchain")
		}

		// Marshal asset back to JSON format
		rentJSON, nerr := json.Marshal(rentAsset)
		if nerr != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}

		return rentJSON, nil
	},
}
