package service

import (
	"github.com/stretchr/testify/assert"
	"order-test/entity"
	"testing"
)

func Test_label_free_shipping_when_value_greater_than_thousand(t *testing.T) {
	product := entity.Product{
		Category: entity.CategoryElectronic,
		Value:    1200,
	}

	orderInput := CreateOrderDto{
		Product: product,
		Method:  entity.PaymentMethodPix,
	}

	result := OrderCreate(orderInput)

	assert.Equal(t, entity.LabelFreeShipping, result.Labels[0])
	assert.Equal(t, product.Value, result.Payment.Value)
}

func Test_label_fragile_when_category_is_home_appliance(t *testing.T) {
	product := entity.Product{
		Category: entity.CategoryHomeAppliance,
		Value:    900,
	}

	orderInput := CreateOrderDto{
		Product: product,
		Method:  entity.PaymentMethodPix,
	}

	result := OrderCreate(orderInput)

	assert.Equal(t, entity.LabelFragile, result.Labels[0])
	assert.Equal(t, product.Value, result.Payment.Value)
}

func Test_label_gift_when_category_is_children(t *testing.T) {
	product := entity.Product{
		Category: entity.CategoryChildren,
		Value:    900,
	}

	orderInput := CreateOrderDto{
		Product: product,
		Method:  entity.PaymentMethodPix,
	}

	result := OrderCreate(orderInput)

	assert.Equal(t, entity.LabelGift, result.Labels[0])
	assert.Equal(t, product.Value, result.Payment.Value)
}

func Test_discount_by_payment_method(t *testing.T) {
	product := entity.Product{
		Category: entity.CategoryChildren,
		Value:    1000,
	}

	orderInput := CreateOrderDto{
		Product: product,
		Method:  entity.PaymentMethodBankSlip,
	}

	expectedValue := 900
	result := OrderCreate(orderInput)

	assert.Equal(t, expectedValue, result.Payment.Value)
}

func Test_multiple_rules(t *testing.T) {
	product := entity.Product{
		Category: entity.CategoryChildren,
		Value:    1100,
	}

	orderInput := CreateOrderDto{
		Product: product,
		Method:  entity.PaymentMethodBankSlip,
	}

	result := OrderCreate(orderInput)
	expectedLabels := []string{entity.LabelFreeShipping, entity.LabelGift}
	expectedValue := 990

	assert.Equal(t, expectedLabels, result.Labels)
	assert.Equal(t, expectedValue, result.Payment.Value)
}
