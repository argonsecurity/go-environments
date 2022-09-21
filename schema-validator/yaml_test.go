package schemavalidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateYaml(t *testing.T) {
	type args struct {
		yamlFilePath   string
		schemaFilePath string
	}
	testCases := []struct {
		name          string
		args          args
		expectedError error
	}{
		{
			name: "yaml is empty",
			args: args{
				yamlFilePath:   "testdata/yamls/empty.yml",
				schemaFilePath: "testdata/schemas/valid.schema.json",
			},
			expectedError: emptyYamlError,
		},
		{
			name: "schema is empty",
			args: args{
				yamlFilePath:   "testdata/yamls/valid.yml",
				schemaFilePath: "testdata/schemas/empty.schema.json",
			},
			expectedError: emptySchemaError,
		},
		{
			name: "yaml is not valid (invalid yaml)",
			args: args{
				yamlFilePath:   "testdata/yamls/invalid.yml",
				schemaFilePath: "testdata/schemas/valid.schema.json",
			},
			expectedError: NewErrParseInput("yaml: line 6: did not find expected node content"),
		},
		{
			name: "yaml does not fit the schema",
			args: args{
				yamlFilePath:   "testdata/yamls/not_fit_to_schema.yml",
				schemaFilePath: "testdata/schemas/valid.schema.json",
			},
			expectedError: NewErrFailedValidation("image: Invalid type. Expected: object, given: string\nservices: Invalid type. Expected: object, given: array\n"),
		},
		{
			name: "yaml fits the schema",
			args: args{
				yamlFilePath:   "testdata/yamls/valid.yml",
				schemaFilePath: "testdata/schemas/valid.schema.json",
			},
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			yamlData := readFile(tc.args.yamlFilePath)
			schemaData := readFile(tc.args.schemaFilePath)

			err := ValidateYaml(yamlData, schemaData)
			assert.Equal(t, tc.expectedError, err, tc.name)
		})
	}

}
