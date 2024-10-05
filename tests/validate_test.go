package tests

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/validation"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Param struct {
	fields contracts.Fields
}

func (p Param) ToFields() contracts.Fields {
	return p.fields
}

func (p Param) Rules() contracts.Fields {
	return contracts.Fields{
		"id":   "required,gte=1",
		"name": "required",
	}
}

func TestValidate(t *testing.T) {

	assert.True(t, len(validation.Valid(contracts.Fields{}, contracts.Fields{
		"name": "required",
	})) == 1)

	assert.True(t, len(validation.Valid(contracts.Fields{
		"name": "啦啦啦",
	}, contracts.Fields{
		"name": "required",
	})) == 0)

	assert.True(t, len(validation.Form(Param{contracts.Fields{
		"id":   1,
		"name": "goal",
	}})) == 0)

}
