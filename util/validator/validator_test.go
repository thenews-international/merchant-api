package validator_test

import (
	"testing"

	"merchant/util/validator"
)

type testCase struct {
	name     string
	input    interface{}
	expected string
}

var tests = []*testCase{
	// --- invalid cases ---
	{
		name: `required`,
		input: struct {
			Email string `json:"email" form:"required"`
		}{},
		expected: "email is a required field",
	},
	{
		name: `min`,
		input: struct {
			Age int `json:"age" form:"min=16"`
		}{Age: 7},
		expected: "age must be a minimum of 16 in length",
	},
	{
		name: `max`,
		input: struct {
			Course string `json:"course" form:"max=7"`
		}{Course: "CS-0001."},
		expected: "course must be a maximum of 7 in length",
	},
	{
		name: `url`,
		input: struct {
			Host string `json:"host" form:"url"`
		}{Host: "foobar.com"},
		expected: "host must be a valid URL",
	},
	{
		name: `eqfield`,
		input: struct {
			Password        string `json:"password"`
			ConfirmPassword string `json:"confirm_password" form:"eqfield=Password"`
		}{
			Password:        "abc123",
			ConfirmPassword: "abc213",
		},
		expected: "confirm_password must be equal to Password",
	},
	{
		name: `uuid4`,
		input: struct {
			Id string `json:"id" form:"uuid4_rfc4122"`
		}{
			Id: "00000000-0000-0000-0000-000000000000",
		},
		expected: "id must be a valid uuid",
	},
	{
		name: `uuid4 with uuid1`,
		input: struct {
			Id string `json:"id" form:"uuid4_rfc4122"`
		}{
			Id: "9c9195c8-1016-4c9a-8cde-48554e888ca",
		},
		expected: "id must be a valid uuid",
	},
	{
		name: `oneof`,
		input: struct {
			Day string `json:"input" form:"required,oneof=saturday sunday"`
		}{
			Day: "monday",
		},
		expected: "input must be one of saturday sunday",
	},
	{name: `date`,
		input: struct {
			Date string `json:"date" form:"datetime=2006-01-02"`
		}{
			Date: "2021-02-29",
		},
		expected: "date must be a valid date",
	},
	{
		name: `month`,
		input: struct {
			Month string `json:"month" form:"datetime=2006-01"`
		}{
			Month: "2021-14",
		},
		expected: "month must follow 2006-01 format",
	},
}

func TestToErrResponse(t *testing.T) {
	vr := validator.New()

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := vr.Struct(tc.input)
			if errResp := validator.ToErrResponse(err); errResp == nil || len(errResp.Errors) != 1 {
				t.Fatalf(`Expected:"{[%v]}", Got:"%v"`, tc.expected, errResp)
			} else if errResp.Errors[0] != tc.expected {
				t.Fatalf(`Expected:"%v", Got:"%v"`, tc.expected, errResp.Errors[0])
			}
		})
	}
}
