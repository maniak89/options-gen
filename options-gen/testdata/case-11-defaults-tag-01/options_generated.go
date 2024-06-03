// Code generated by options-gen. DO NOT EDIT.
package testcase

import (
	"fmt"
	fmt461e464ebed9 "fmt"
	"io"
	"time"

	errors461e464ebed9 "github.com/kazhuravlev/options-gen/pkg/errors"
	validator461e464ebed9 "github.com/kazhuravlev/options-gen/pkg/validator"
)

type optField int8

const (
	FieldrWCloser    optField = 0
	FieldoptStringer optField = 1
	FieldvalInt      optField = 2
	FieldvalInt8     optField = 3
	FieldvalInt16    optField = 4
	FieldvalInt32    optField = 5
	FieldvalInt64    optField = 6
	FieldvalUInt     optField = 7
	FieldvalUInt8    optField = 8
	FieldvalUInt16   optField = 9
	FieldvalUInt32   optField = 10
	FieldvalUInt64   optField = 11
	FieldvalFloat32  optField = 12
	FieldvalFloat64  optField = 13
	FieldvalDuration optField = 14
	FieldvalString   optField = 15
	FieldvalBool     optField = 16
)

var optIsSet = [17]bool{}

type OptOptionsSetter func(o *Options)

func NewOptions(
	rWCloser io.ReadWriteCloser,
	options ...OptOptionsSetter,
) Options {
	o := Options{}

	var empty [17]bool
	optIsSet = empty

	// Setting defaults from field tag (if present)
	o.valInt = 1
	optIsSet[FieldvalInt] = true
	o.valInt8 = 8
	optIsSet[FieldvalInt8] = true
	o.valInt16 = 16
	optIsSet[FieldvalInt16] = true
	o.valInt32 = 32
	optIsSet[FieldvalInt32] = true
	o.valInt64 = 64
	optIsSet[FieldvalInt64] = true
	o.valUInt = 11
	optIsSet[FieldvalUInt] = true
	o.valUInt8 = 88
	optIsSet[FieldvalUInt8] = true
	o.valUInt16 = 1616
	optIsSet[FieldvalUInt16] = true
	o.valUInt32 = 3232
	optIsSet[FieldvalUInt32] = true
	o.valUInt64 = 6464
	optIsSet[FieldvalUInt64] = true
	o.valFloat32 = 32.32
	optIsSet[FieldvalFloat32] = true
	o.valFloat64 = 64.64
	optIsSet[FieldvalFloat64] = true
	o.valDuration, _ = time.ParseDuration("3s")
	optIsSet[FieldvalDuration] = true
	o.valString = "golang"
	optIsSet[FieldvalString] = true
	o.valBool = true
	optIsSet[FieldvalBool] = true

	o.rWCloser = rWCloser
	optIsSet[FieldrWCloser] = true

	for _, opt := range options {
		opt(&o)
	}
	return o
}

func WithOptStringer(opt fmt.Stringer) OptOptionsSetter {
	return func(o *Options) {
		o.optStringer = opt
		optIsSet[FieldoptStringer] = true
	}
}

func WithValInt(opt int) OptOptionsSetter {
	return func(o *Options) {
		o.valInt = opt
		optIsSet[FieldvalInt] = true
	}
}

func WithValInt8(opt int8) OptOptionsSetter {
	return func(o *Options) {
		o.valInt8 = opt
		optIsSet[FieldvalInt8] = true
	}
}

func WithValInt16(opt int16) OptOptionsSetter {
	return func(o *Options) {
		o.valInt16 = opt
		optIsSet[FieldvalInt16] = true
	}
}

func WithValInt32(opt int32) OptOptionsSetter {
	return func(o *Options) {
		o.valInt32 = opt
		optIsSet[FieldvalInt32] = true
	}
}

func WithValInt64(opt int64) OptOptionsSetter {
	return func(o *Options) {
		o.valInt64 = opt
		optIsSet[FieldvalInt64] = true
	}
}

func WithValUInt(opt uint) OptOptionsSetter {
	return func(o *Options) {
		o.valUInt = opt
		optIsSet[FieldvalUInt] = true
	}
}

func WithValUInt8(opt uint8) OptOptionsSetter {
	return func(o *Options) {
		o.valUInt8 = opt
		optIsSet[FieldvalUInt8] = true
	}
}

func WithValUInt16(opt uint16) OptOptionsSetter {
	return func(o *Options) {
		o.valUInt16 = opt
		optIsSet[FieldvalUInt16] = true
	}
}

func WithValUInt32(opt uint32) OptOptionsSetter {
	return func(o *Options) {
		o.valUInt32 = opt
		optIsSet[FieldvalUInt32] = true
	}
}

func WithValUInt64(opt uint64) OptOptionsSetter {
	return func(o *Options) {
		o.valUInt64 = opt
		optIsSet[FieldvalUInt64] = true
	}
}

func WithValFloat32(opt float32) OptOptionsSetter {
	return func(o *Options) {
		o.valFloat32 = opt
		optIsSet[FieldvalFloat32] = true
	}
}

func WithValFloat64(opt float64) OptOptionsSetter {
	return func(o *Options) {
		o.valFloat64 = opt
		optIsSet[FieldvalFloat64] = true
	}
}

func WithValDuration(opt time.Duration) OptOptionsSetter {
	return func(o *Options) {
		o.valDuration = opt
		optIsSet[FieldvalDuration] = true
	}
}

func WithValString(opt string) OptOptionsSetter {
	return func(o *Options) {
		o.valString = opt
		optIsSet[FieldvalString] = true
	}
}

func WithValBool(opt bool) OptOptionsSetter {
	return func(o *Options) {
		o.valBool = opt
		optIsSet[FieldvalBool] = true
	}
}

func (o *Options) Validate() error {
	errs := new(errors461e464ebed9.ValidationErrors)
	errs.Add(errors461e464ebed9.NewValidationError("valUInt8", _validate_Options_valUInt8(o)))
	errs.Add(errors461e464ebed9.NewValidationError("valDuration", _validate_Options_valDuration(o)))
	errs.Add(errors461e464ebed9.NewValidationError("valString", _validate_Options_valString(o)))
	return errs.AsError()
}

func (o *Options) IsSet(field optField) bool {
	return optIsSet[field]
}

func _validate_Options_valUInt8(o *Options) error {
	if err := validator461e464ebed9.GetValidatorFor(o).Var(o.valUInt8, "min=50"); err != nil {
		return fmt461e464ebed9.Errorf("field `valUInt8` did not pass the test: %w", err)
	}
	return nil
}

func _validate_Options_valDuration(o *Options) error {
	if err := validator461e464ebed9.GetValidatorFor(o).Var(o.valDuration, "min=100ms,max=30s"); err != nil {
		return fmt461e464ebed9.Errorf("field `valDuration` did not pass the test: %w", err)
	}
	return nil
}

func _validate_Options_valString(o *Options) error {
	if err := validator461e464ebed9.GetValidatorFor(o).Var(o.valString, "required"); err != nil {
		return fmt461e464ebed9.Errorf("field `valString` did not pass the test: %w", err)
	}
	return nil
}
