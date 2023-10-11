package usecase

import (
	"order-test/entity"
	"order-test/rules"
)

const freeShippingValue = 1000
const discountPercentage = 0.10

type CreateOrderDto struct {
	Product entity.Product
	Method  string
}

func OrderCreate(orderDto CreateOrderDto) entity.Order {

	payment := entity.Payment{
		Method: orderDto.Method,
		Value:  orderDto.Product.Value,
	}

	order := entity.Order{
		Product: orderDto.Product,
		Payment: payment,
	}

	rulesOrder := rules.Rule{
		Order: &order,
	}

	rulesOrder.LabelByValue(freeShippingValue, entity.LabelFreeShipping)
	rulesOrder.LabelByCategory(entity.CategoryHomeAppliance, entity.LabelFragile)
	rulesOrder.LabelByCategory(entity.CategoryChildren, entity.LabelGift)
	rulesOrder.DiscountByPaymentMethod(entity.PaymentMethodBankSlip, discountPercentage)

	return order
}
