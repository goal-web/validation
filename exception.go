package validation

import (
	"errors"
	"github.com/goal-web/contracts"
)

type Exception struct {
	Err    error
	Param  contracts.Fields
	Errors contracts.Fields
}

func (e *Exception) Error() string {
	return e.Err.Error()
}

func (e *Exception) GetPrevious() contracts.Exception {
	return nil
}

func NewException(param contracts.Fields, errs contracts.Fields) contracts.Exception {
	return &Exception{Err: errors.New("param validation failed"), Param: param, Errors: errs}
}
