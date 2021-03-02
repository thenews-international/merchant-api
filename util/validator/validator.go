package validator

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ErrResponse struct {
	Errors []string `json:"errors"`
}

func New() *validator.Validate {
	validate := validator.New()
	validate.SetTagName("form")

	// Using the names which have been specified for JSON representations of structs, rather than normal Go field names
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return validate
}

func ToErrResponse(err error) *ErrResponse {
	if fieldErrors, ok := err.(validator.ValidationErrors); ok {
		resp := ErrResponse{
			Errors: make([]string, len(fieldErrors)),
		}

		for i, err := range fieldErrors {
			switch err.Tag() {
			case "required":
				resp.Errors[i] = fmt.Sprintf("%s is a required field", err.Field())
			case "required_if":
				resp.Errors[i] = fmt.Sprintf("%s is a required field", err.Field())
			case "required_with":
				resp.Errors[i] = fmt.Sprintf("%s is a required field", err.Field())
			case "eqfield":
				resp.Errors[i] = fmt.Sprintf("%s must be equal to %s", err.Field(), err.Param())
			case "gt":
				resp.Errors[i] = fmt.Sprintf("%s must be greater than %s", err.Field(), err.Param())
			case "min":
				resp.Errors[i] = fmt.Sprintf("%s must be a minimum of %s in length", err.Field(), err.Param())
			case "max":
				resp.Errors[i] = fmt.Sprintf("%s must be a maximum of %s in length", err.Field(), err.Param())
			case "email":
				resp.Errors[i] = fmt.Sprintf("%s must be a valid email address", err.Field())
			case "url":
				resp.Errors[i] = fmt.Sprintf("%s must be a valid URL", err.Field())
			case "uuid4_rfc4122":
				resp.Errors[i] = fmt.Sprintf("%s must be a valid uuid", err.Field())
			case "oneof":
				resp.Errors[i] = fmt.Sprintf("%s must be one of %s", err.Field(), err.Param())
			case "datetime":
				if err.Param() == "2006-01-02" {
					resp.Errors[i] = fmt.Sprintf("%s must be a valid date", err.Field())
				} else {
					resp.Errors[i] = fmt.Sprintf("%s must follow %s format", err.Field(), err.Param())
				}
			default:
				resp.Errors[i] = fmt.Sprintf("something wrong on %s; %s", err.Field(), err.Tag())
			}
		}

		return &resp
	}

	return nil
}
