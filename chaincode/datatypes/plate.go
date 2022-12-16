package datatypes

import (
	"strings"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
)

var plate = assets.DataType{
	Parse: func(data interface{}) (string, interface{}, errors.ICCError) {
		plate, ok := data.(string)

		plate = strings.ReplaceAll(plate, " ", "")
		plate = strings.ReplaceAll(plate, "-", "")
		plate = strings.ReplaceAll(plate, ".", "")
		plate = strings.ReplaceAll(plate, "/", "")
		plate = strings.ToUpper(plate)

		if !ok {
			return "", nil, errors.NewCCError("property must be a string", 400)
		}

		if len(plate) != 6 {
			return "", nil, errors.NewCCError("plate must have 6 characters", 400)
		}

		return plate, plate, nil
	},
}
