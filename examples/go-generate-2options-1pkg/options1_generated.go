// Code generated by options-gen. DO NOT EDIT.
package gogenerate

import (
	fmt461e464ebed9 "fmt"

	errors461e464ebed9 "github.com/kazhuravlev/options-gen/pkg/errors"
	validator461e464ebed9 "github.com/kazhuravlev/options-gen/pkg/validator"
)

type OptOptions1Setter func(o *Options1)

func NewOptions1(
	options ...OptOptions1Setter,
) Options1 {
	o := Options1{}

	// Setting defaults from variable
	o.field0 = defaultOptions1.field0

	o.field1 = defaultOptions1.field1

	o.field2 = defaultOptions1.field2

	o.field3 = defaultOptions1.field3

	for _, opt := range options {
		opt(&o)
	}
	return o
}

// Options1.field0
func WithKKKField0(opt int) OptOptions1Setter {
	return func(o *Options1) {
		o.field0 = opt

	}
}

// Options1.field1
func WithKKKField1(opt int) OptOptions1Setter {
	return func(o *Options1) {
		o.field1 = opt

	}
}

// Options1.field2
func WithKKKField2(opt int) OptOptions1Setter {
	return func(o *Options1) {
		o.field2 = opt

	}
}

// Options1.field3
func WithKKKField3(opt int) OptOptions1Setter {
	return func(o *Options1) {
		o.field3 = opt

	}
}

func (o *Options1) Validate() error {
	errs := new(errors461e464ebed9.ValidationErrors)
	errs.Add(errors461e464ebed9.NewValidationError("field0", _validate_Options1_field0(o)))
	errs.Add(errors461e464ebed9.NewValidationError("field1", _validate_Options1_field1(o)))
	errs.Add(errors461e464ebed9.NewValidationError("field2", _validate_Options1_field2(o)))
	errs.Add(errors461e464ebed9.NewValidationError("field3", _validate_Options1_field3(o)))
	return errs.AsError()
}

func _validate_Options1_field0(o *Options1) error {
	if err := validator461e464ebed9.GetValidatorFor(o).Var(o.field0, "min:3"); err != nil {
		return fmt461e464ebed9.Errorf("field `field0` did not pass the test: %w", err)
	}
	return nil
}

func _validate_Options1_field1(o *Options1) error {
	if err := validator461e464ebed9.GetValidatorFor(o).Var(o.field1, "min:3"); err != nil {
		return fmt461e464ebed9.Errorf("field `field1` did not pass the test: %w", err)
	}
	return nil
}

func _validate_Options1_field2(o *Options1) error {
	if err := validator461e464ebed9.GetValidatorFor(o).Var(o.field2, "min:3"); err != nil {
		return fmt461e464ebed9.Errorf("field `field2` did not pass the test: %w", err)
	}
	return nil
}

func _validate_Options1_field3(o *Options1) error {
	if err := validator461e464ebed9.GetValidatorFor(o).Var(o.field3, "min:3"); err != nil {
		return fmt461e464ebed9.Errorf("field `field3` did not pass the test: %w", err)
	}
	return nil
}
