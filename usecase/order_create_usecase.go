package usecase

import (
	"order-test/entity"
	"order-test/rules"
)

const freeShippingValue = 1000

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

	rules.LabelByValue(freeShippingValue, entity.LabelFreeShipping, &order)
	rules.LabelByCategory(entity.CategoryHomeAppliance, entity.LabelFragile, &order)
	rules.LabelByCategory(entity.CategoryChildren, entity.LabelGift, &order)
	rules.DiscountByPaymentMethod(entity.PaymentMethodBankSlip, 0.10, &order)

	return order
}
