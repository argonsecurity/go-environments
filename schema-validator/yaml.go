package schemavalidator

import (
	"gopkg.in/yaml.v3"
)

func ValidateYaml(yamlData []byte, schemaData []byte) error {
	if len(yamlData) == 0 {
		return emptyYamlError
	}

	if len(schemaData) == 0 {
		return emptySchemaError
	}

	var m map[string]any
	if err := yaml.Unmarshal(yamlData, &m); err != nil {
		return NewErrParseInput(err.Error())
	}

	return validate(m, schemaData)
}
