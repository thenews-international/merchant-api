package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"merchant/model"
)

var merchantDto = &model.MerchantDto{
	ID:           "8336fc00-43b5-40f7-83e3-27c018058054",
	Email:        "admin@pacenow.com",
	BusinessName: "PaceNow",
	Description:  "",
	Status:       "Active",
}
var merchant = &model.Merchant{
	Model: model.Model{
		ID: "8336fc00-43b5-40f7-83e3-27c018058054",
	},
	Email:        "admin@pacenow.com",
	BusinessName: "PaceNow",
	Description:  "",
	Status:       "Active",
}

var merchantDtos = model.MerchantDtos{
	merchantDto,
}

var merchants = model.Merchants{
	merchant,
}

func TestMerchantToDto(t *testing.T) {
	t.Parallel()

	expected := merchantDto
	actual := merchant.ToDto()
	assert.Equalf(t, expected, actual, "Expected: %q, Actual: %q", expected, actual)
}

func TestCountriesToDto(t *testing.T) {
	t.Parallel()

	expected := merchantDtos
	actual := merchants.ToDto()
	assert.Equalf(t, expected, actual, "Expected: %q, Actual: %q", expected, actual)
}
