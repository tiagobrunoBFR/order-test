package usecase

import (
	"github.com/stretchr/testify/assert"
	"order-test/entity"
	"testing"
)

func TestOrderCreate(t *testing.T) {
	casesTest := []struct {
		productCategory      string
		productValue         int
		paymentMethod        string
		expectedLabel        string
		expectedPaymentValue int
	}{
		{entity.CategoryElectronic, 1200, entity.PaymentMethodPix, entity.LabelFreeShipping, 1200},
		{entity.CategoryHomeAppliance, 900, entity.PaymentMethodPix, entity.LabelFragile, 900},
		{entity.CategoryChildren, 900, entity.PaymentMethodPix, entity.LabelGift, 900},
		{entity.CategoryElectronic, 900, entity.PaymentMethodBankSlip, "", 900},
	}

	for _, caseTest := range casesTest {
		product := entity.Product{
			Category: caseTest.productCategory,
			Value:    caseTest.productValue,
		}

		orderInput := CreateOrderDto{
			Product: product,
			Method:  caseTest.paymentMethod,
		}

		result := OrderCreate(orderInput)
		var resultLabel string
		if len(result.Labels) > 0 {
			resultLabel = result.Labels[0]
		}

		assert.Equal(t, caseTest.expectedLabel, resultLabel)
		assert.Equal(t, caseTest.expectedPaymentValue, result.Payment.Value)
	}
}
