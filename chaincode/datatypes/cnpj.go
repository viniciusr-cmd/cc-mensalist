package datatypes

import (
	"strings"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
)

var cnpj = assets.DataType{
	AcceptedFormats: []string{"string"},
	Parse: func(data interface{}) (string, interface{}, errors.ICCError) {
		cnpj, ok := data.(string)
		if !ok {
			return "", nil, errors.NewCCError("property must be a string", 400)
		}

		cnpj = strings.ReplaceAll(cnpj, ".", "")
		cnpj = strings.ReplaceAll(cnpj, "-", "")
		cnpj = strings.ReplaceAll(cnpj, "/", "")
		cnpj = strings.ReplaceAll(cnpj, " ", "")

		if len(cnpj) != 14 {
			return "", nil, errors.NewCCError("CNPJ must have 14 digits", 400)
		}

		// CNPJ must not have any letters
		if strings.ContainsAny(cnpj, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
			return "", nil, errors.NewCCError("CNPJ must not have letters", 400)
		}

		//CNPJ must not have any repeated digits
		if strings.Repeat(cnpj[0:1], 14) == cnpj {
			return "", nil, errors.NewCCError("CNPJ must not have repeated digits", 400)
		}

		// CNPJ Validation
		// Validate first digit
		var vd0 int
		for i, d := range cnpj {
			if i >= 12 {
				break
			}
			dnum := int(d) - '0'

			if i < 4 {
				vd0 += (5 - i) * dnum
			} else {
				vd0 += (13 - i) * dnum
			}
		}

		vd0 = 11 - vd0%11
		if vd0 > 9 {
			vd0 = 0
		}
		if int(cnpj[12])-'0' != vd0 {
			return "", nil, errors.NewCCError("Invalid CNPJ", 400)
		}

		// Validate second digit
		var vd1 int
		for i, d := range cnpj {
			if i >= 13 {
				break
			}
			dnum := int(d) - '0'

			if i < 5 {
				vd1 += (6 - i) * dnum
			} else {
				vd1 += (14 - i) * dnum
			}
		}
		vd1 = 11 - vd1%11
		if vd1 > 9 {
			vd1 = 0
		}
		if int(cnpj[13])-'0' != vd1 {
			return "", nil, errors.NewCCError("Invalid CNPJ", 400)
		}

		return cnpj, cnpj, nil
	},
}
