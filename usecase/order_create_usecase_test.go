package usecase

import (
	"github.com/stretchr/testify/assert"
	"order-test/entity"
	"testing"
)

func TestOrderCreate(t *testing.T) {
	product := entity.Product{
		Category: "eletrodoméstico",
		Value:    1200,
	}

	method := "pix"

	expectedLabelFreeShipping := entity.FreeShippingLabel
	expectedLabelFragile := entity.FragileLabel

	orderInput := CreateOrderDto{
		Product: product,
		Method:  method,
	}

	result := OrderCreate(orderInput)
	assert.Equal(t, expectedLabelFreeShipping, result.Labels[0])
	assert.Equal(t, expectedLabelFragile, result.Labels[1])
}
