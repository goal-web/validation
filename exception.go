package validation

import (
	"github.com/goal-web/contracts"
)

type Exception struct {
	param  contracts.Fields
	errors contracts.Fields
	string
}

func NewException(param contracts.Fields, errors contracts.Fields) Exception {
	return Exception{param, errors, "param validation failed"}
}

func (this Exception) Error() string {
	return this.string
}

func (this Exception) Fields() contracts.Fields {
	return this.param
}

func (this Exception) GetErrors() contracts.Fields {
	return this.errors
}
