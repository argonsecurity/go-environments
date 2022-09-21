package schemavalidator

import (
	"fmt"
	"os"

	"github.com/xeipuuv/gojsonschema"
	"golang.org/x/exp/slices"
)

func validate(object map[string]any, schemaData []byte) error {
	schemaLoader := gojsonschema.NewBytesLoader(schemaData)
	documentLoader := gojsonschema.NewGoLoader(object)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)

	if err != nil {
		return NewErrParseSchema(err.Error())
	}

	if !result.Valid() {
		return NewErrFailedValidation(parseValidationErrorResults(result.Errors()))
	}
	return nil
}

func parseValidationErrorResults(errorResults []gojsonschema.ResultError) string {
	slices.SortFunc(errorResults, func(a, b gojsonschema.ResultError) bool {
		return a.String() < b.String()
	})

	var error string
	for _, result := range errorResults {
		error = fmt.Sprintf("%s%s\n", error, result.String())
	}
	return error
}

func readFile(filename string) []byte {
	b, _ := os.ReadFile(filename)
	return b
}
