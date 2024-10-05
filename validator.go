package validation

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
)

var Validator = validator.New()

func Struct(data interface{}) error {
	return Validator.Struct(data)
}

func Valid(data interface{}, rules contracts.Fields) contracts.Fields {
	switch param := data.(type) {
	case contracts.Fields:
		return Validator.ValidateMap(param, rules)
	case contracts.FieldsProvider:
		return Validator.ValidateMap(param.ToFields(), rules)
	}

	fields, err := utils.ToFields(data)
	if err != nil {
		panic(Exception{
			Param:  fields,
			Errors: nil,
			Err:    errors.New("unsupported parameter type"),
		})
	}

	return Validator.ValidateMap(fields, rules)
}

func Form(validatable contracts.Validatable) contracts.Fields {
	return Validator.ValidateMap(validatable.ToFields(), validatable.Rules())
}

func VerifyForm(validatable contracts.Validatable) {
	if errs := Form(validatable); len(errs) > 0 {
		panic(NewException(validatable.ToFields(), validatable.Rules()))
	}
}

func VerifyStruct(data interface{}) {
	if err := Struct(data); err != nil {
		panic(err)
	}
}

func Verify(data interface{}, rules contracts.Fields) {
	if errs := Valid(data, rules); len(errs) > 0 {
		var fields, _ = utils.ToFields(data)
		panic(NewException(fields, rules))
	}
}
