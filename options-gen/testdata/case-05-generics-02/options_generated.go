// Code generated by options-gen. DO NOT EDIT.

package testcase

import (
	fmt461e464ebed9 "fmt"
	"net/http"

	errors461e464ebed9 "github.com/kazhuravlev/options-gen/pkg/errors"
	validator461e464ebed9 "github.com/kazhuravlev/options-gen/pkg/validator"
)

type optSomeField int8

const (
	FieldSomerequiredHandler optSomeField = 0
	FieldSomerequiredKey     optSomeField = 1
	FieldSomehandler         optSomeField = 2
	FieldSomekey             optSomeField = 3
	FieldSomeoptHandler      optSomeField = 4
	FieldSomeoptKey          optSomeField = 5
	FieldSomeanyOpt          optSomeField = 6
)

var optSomeIsSet = [7]bool{}

type OptOptionsSetter[KeyT int | string, TT any] func(o *Options[KeyT, TT])

func NewOptions[KeyT int | string, TT any](
	requiredHandler http.Handler,
	requiredKey KeyT,
	handler http.Handler,
	key KeyT,
	options ...OptOptionsSetter[KeyT, TT],
) Options[KeyT, TT] {
	o := Options[KeyT, TT]{}

	var empty [7]bool
	optSomeIsSet = empty

	// Setting defaults from field tag (if present)

	o.requiredHandler = requiredHandler
	optSomeIsSet[FieldSomerequiredHandler] = true
	o.requiredKey = requiredKey
	optSomeIsSet[FieldSomerequiredKey] = true
	o.handler = handler
	optSomeIsSet[FieldSomehandler] = true
	o.key = key
	optSomeIsSet[FieldSomekey] = true

	for _, opt := range options {
		opt(&o)
	}
	return o
}

func WithSomeOptHandler[KeyT int | string, TT any](opt http.Handler) OptOptionsSetter[KeyT, TT] {
	return func(o *Options[KeyT, TT]) {
		o.optHandler = opt
		optSomeIsSet[FieldSomeoptHandler] = true
	}
}

func WithSomeOptKey[KeyT int | string, TT any](opt KeyT) OptOptionsSetter[KeyT, TT] {
	return func(o *Options[KeyT, TT]) {
		o.optKey = opt
		optSomeIsSet[FieldSomeoptKey] = true
	}
}

func WithSomeAnyOpt[KeyT int | string, TT any](opt TT) OptOptionsSetter[KeyT, TT] {
	return func(o *Options[KeyT, TT]) {
		o.anyOpt = opt
		optSomeIsSet[FieldSomeanyOpt] = true
	}
}

func (o *Options[KeyT, TT]) Validate() error {
	errs := new(errors461e464ebed9.ValidationErrors)
	errs.Add(errors461e464ebed9.NewValidationError("requiredHandler", _validate_Options_requiredHandler[KeyT, TT](o)))
	errs.Add(errors461e464ebed9.NewValidationError("requiredKey", _validate_Options_requiredKey[KeyT, TT](o)))
	return errs.AsError()
}

func (o *Options[KeyT, TT]) IsSet(field optSomeField) bool {
	return optSomeIsSet[field]
}

func _validate_Options_requiredHandler[KeyT int | string, TT any](o *Options[KeyT, TT]) error {
	if err := validator461e464ebed9.GetValidatorFor(o).Var(o.requiredHandler, "required"); err != nil {
		return fmt461e464ebed9.Errorf("field `requiredHandler` did not pass the test: %w", err)
	}
	return nil
}

func _validate_Options_requiredKey[KeyT int | string, TT any](o *Options[KeyT, TT]) error {
	if err := validator461e464ebed9.GetValidatorFor(o).Var(o.requiredKey, "required"); err != nil {
		return fmt461e464ebed9.Errorf("field `requiredKey` did not pass the test: %w", err)
	}
	return nil
}
