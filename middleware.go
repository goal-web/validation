package validation

import "github.com/goal-web/contracts"

func VerifyValidatable(request contracts.Context, pipe contracts.Pipe) interface{} {

	if form, ok := request.(contracts.Validatable); ok {
		VerifyForm(form)
	}

	return pipe(request)
}
