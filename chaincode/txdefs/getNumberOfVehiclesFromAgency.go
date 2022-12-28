package txdefs

import (
	"encoding/json"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	sw "github.com/goledgerdev/cc-tools/stubwrapper"
	tx "github.com/goledgerdev/cc-tools/transactions"
)

// Return the number of Vehicles of a agency
// GET method
var GetNumberOfVehiclesFromAgency = tx.Transaction{
	Tag:         "getNumberOfVehiclesFromAgency",
	Label:       "Get Number Of Vehicles from Agency",
	Description: "Return the number of Vehicles of a agency",
	Method:      "GET",
	Callers:     []string{"$org2MSP", "$orgMSP"}, // Only org2 can call this transactions

	Args: []tx.Argument{
		{
			Tag:         "agency",
			Label:       "Agency",
			Description: "Agency",
			DataType:    "->agency",
			Required:    true,
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		agencyKey, _ := req["agency"].(assets.Key)

		// Returns Agency from channel
		agencyMap, err := agencyKey.GetMap(stub)
		if err != nil {
			return nil, errors.WrapError(err, "failed to get asset from the ledger")
		}

		numberOfVehicles := 0
		vehicles, ok := agencyMap["vehicleList"].([]interface{})
		if ok {
			numberOfVehicles = len(vehicles)
		}

		returnMap := make(map[string]interface{})
		returnMap["numberOfVehicles"] = numberOfVehicles

		// Marshal asset back to JSON format
		returnJSON, nerr := json.Marshal(returnMap)
		if nerr != nil {
			return nil, errors.WrapError(err, "failed to marshal response")
		}

		return returnJSON, nil
	},
}
