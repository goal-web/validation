package validation

import (
	"errors"
	"fmt"
	"github.com/goal-web/contracts"
)

type Exception struct {
	Err    error
	Param  contracts.Fields
	Errors contracts.Fields
}

func (e *Exception) Error() string {
	msg := e.Err.Error()
	for key, value := range e.Errors {
		msg += fmt.Sprintf(". [%s]: %v", key, value)
	}
	return msg
}

func (e *Exception) GetPrevious() contracts.Exception {
	return nil
}

func NewException(param contracts.Fields, errs contracts.Fields) contracts.Exception {
	return &Exception{Err: errors.New("param validation failed"), Param: param, Errors: errs}
}
