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

	labelFreeShipping := rules.LabelByValue{MaxValue: freeShippingValue, Label: entity.LabelFreeShipping}
	labelFragile := rules.LabelByCategory{CategoryExpected: entity.CategoryHomeAppliance, Label: entity.LabelFragile}
	labelGift := rules.LabelByCategory{CategoryExpected: entity.CategoryChildren, Label: entity.LabelGift}
	discountByPayment := rules.DiscountByPaymentMethod{MethodExpected: entity.PaymentMethodBankSlip, DiscountPercentage: 0.10}

	labelFreeShipping.Apply(&order)
	labelFragile.Apply(&order)
	labelGift.Apply(&order)
	discountByPayment.Apply(&order)

	return order
}
