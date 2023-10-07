package usecase

import (
	"github.com/stretchr/testify/assert"
	"order-test/entity"
	"testing"
)

func TestOrderCreate(t *testing.T) {
	product := entity.Product{
		Category: "Test",
		Value:    1200,
	}

	method := "pix"

	expectedLabel := entity.FreeShippingLabel

	orderInput := CreateOrderDto{
		Product: product,
		Method:  method,
	}

	result := OrderCreate(orderInput)
	assert.Equal(t, expectedLabel, result.Labels[0])
}
