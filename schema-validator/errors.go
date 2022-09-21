package schemavalidator

import "fmt"

var (
	emptyYamlError   = NewErrParseInput("yaml is empty")
	emptyJsonError   = NewErrParseInput("json is empty")
	emptySchemaError = NewErrParseSchema("schema is empty")
)

type ErrParseInput struct {
	Message string
}

func (e *ErrParseInput) Error() string {
	return fmt.Sprintf("failed to parse input - %s", e.Message)
}

func NewErrParseInput(message string) *ErrParseInput {
	return &ErrParseInput{
		Message: message,
	}
}

type ErrParseSchema struct {
	Message string
}

func (e *ErrParseSchema) Error() string {
	return fmt.Sprintf("failed to parse schema - %s", e.Message)
}

func NewErrParseSchema(message string) *ErrParseSchema {
	return &ErrParseSchema{
		Message: message,
	}
}

type ErrFailedValidation struct {
	Message string
}

func (e *ErrFailedValidation) Error() string {
	return fmt.Sprintf("schema validation failed - %s", e.Message)
}

func NewErrFailedValidation(message string) *ErrFailedValidation {
	return &ErrFailedValidation{
		Message: message,
	}
}
