package generator //nolint:testpackage

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_checkDefaultValue_Negative(t *testing.T) {
	cases := []struct {
		t   string
		val string
	}{
		{t: "int", val: "a"},
		{t: "int8", val: "b"},
		{t: "int16", val: "c"},
		{t: "int32", val: "d"},
		{t: "int64", val: "e"},

		{t: "uint", val: "aa"},
		{t: "uint8", val: "bb"},
		{t: "uint16", val: "cc"},
		{t: "uint32", val: "dd"},
		{t: "uint64", val: "ee"},

		{t: "float32", val: "aaa"},
		{t: "float64", val: "bbb"},

		{t: "bool", val: "a"},

		{t: "time.Duration", val: "1year"},

		{t: "fmt.Stringer", val: "nil"},
		{t: "Number", val: "nil"},
		{t: "localIterface", val: "nil"},
		{t: "*T", val: "nil"},
	}

	for _, tt := range cases {
		t.Run(tt.t, func(t *testing.T) {
			_, err := checkDefaultValue(tt.t, tt.val)
			assert.Error(t, err)
		})
	}
}

func Test_checkDefaultValue(t *testing.T) {
	cases := []struct {
		t        string
		val      string
		expected string
	}{
		{t: "int", val: "1", expected: "1"},
		{t: "int", val: "-1", expected: "-1"},
		{t: "int8", val: "1", expected: "1"},
		{t: "int8", val: "-1", expected: "-1"},
		{t: "int16", val: "1", expected: "1"},
		{t: "int16", val: "-1", expected: "-1"},
		{t: "int32", val: "1", expected: "1"},
		{t: "int32", val: "-1", expected: "-1"},
		{t: "int64", val: "1", expected: "1"},
		{t: "int64", val: "-1", expected: "-1"},

		{t: "uint", val: "1", expected: "1"},
		{t: "uint8", val: "1", expected: "1"},
		{t: "uint16", val: "1", expected: "1"},
		{t: "uint32", val: "1", expected: "1"},
		{t: "uint64", val: "1", expected: "1"},

		{t: "float32", val: "3.14", expected: "3.14"},
		{t: "float32", val: "-3.14", expected: "-3.14"},
		{t: "float64", val: "3.14", expected: "3.14"},
		{t: "float64", val: "-3.14", expected: "-3.14"},

		{t: "bool", val: "1", expected: "true"},
		{t: "bool", val: "t", expected: "true"},
		{t: "bool", val: "T", expected: "true"},
		{t: "bool", val: "true", expected: "true"},
		{t: "bool", val: "TRUE", expected: "true"},
		{t: "bool", val: "True", expected: "true"},
		{t: "bool", val: "0", expected: "false"},
		{t: "bool", val: "f", expected: "false"},
		{t: "bool", val: "F", expected: "false"},
		{t: "bool", val: "false", expected: "false"},
		{t: "bool", val: "FALSE", expected: "false"},
		{t: "bool", val: "False", expected: "false"},

		{t: "time.Duration", val: "1h", expected: "1h"},

		// {t: "fmt.Stringer", val: "nil"},
		// {t: "Number", val: "nil"},
		// {t: "localIterface", val: "nil"},
		// {t: "*T", val: "nil"},
	}

	for _, tt := range cases {
		t.Run(tt.t, func(t *testing.T) {
			v, err := checkDefaultValue(tt.t, tt.val)
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, fmt.Sprintf("%v", v))
		})
	}
}
