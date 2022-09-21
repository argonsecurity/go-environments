package schemavalidator

import (
	"encoding/json"
)

func ValidateJson(jsonData []byte, schemaData []byte) error {
	if len(jsonData) == 0 {
		return emptyJsonError
	}

	if len(schemaData) == 0 {
		return emptySchemaError
	}

	var m map[string]any
	if err := json.Unmarshal(jsonData, &m); err != nil {
		return NewErrParseInput(err.Error())
	}

	return validate(m, schemaData)
}
