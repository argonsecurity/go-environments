package schemavalidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateJson(t *testing.T) {
	type args struct {
		jsonFilePath   string
		schemaFilePath string
	}
	tests := []struct {
		name          string
		args          args
		expectedError error
	}{
		{
			name: "json is empty",
			args: args{
				jsonFilePath:   "testdata/jsons/empty.json",
				schemaFilePath: "testdata/schemas/valid.schema.json",
			},
			expectedError: emptyJsonError,
		},
		{
			name: "schema is empty",
			args: args{
				jsonFilePath:   "testdata/jsons/valid.json",
				schemaFilePath: "testdata/schemas/empty.schema.json",
			},
			expectedError: emptySchemaError,
		},
		{
			name: "json is not valid (invalid json)",
			args: args{
				jsonFilePath:   "testdata/jsons/invalid.json",
				schemaFilePath: "testdata/schemas/valid.schema.json",
			},
			expectedError: NewErrParseInput("invalid character '}' looking for beginning of value"),
		},
		{
			name: "json does not fit the schema",
			args: args{
				jsonFilePath:   "testdata/jsons/not_fit_to_schema.json",
				schemaFilePath: "testdata/schemas/valid.schema.json",
			},
			expectedError: NewErrFailedValidation("image: Invalid type. Expected: object, given: string\n"),
		},
		{
			name: "json fits the schema",
			args: args{
				jsonFilePath:   "testdata/jsons/valid.json",
				schemaFilePath: "testdata/schemas/valid.schema.json",
			},
			expectedError: nil,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			jsonData := readFile(tc.args.jsonFilePath)
			schemaData := readFile(tc.args.schemaFilePath)

			err := ValidateJson(jsonData, schemaData)
			assert.Equal(t, tc.expectedError, err, tc.name)
		})
	}
}
