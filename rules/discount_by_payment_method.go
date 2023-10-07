package rules

import "order-test/entity"

const discountPercentage = 0.10

func DiscountByPaymentMethod(method string, o *entity.Order) {
	if method == entity.PaymentMethodBankSlip {
		var value float64
		value = float64(o.Payment.Value / 100)
		value = value * discountPercentage
		o.Payment.Value = o.Payment.Value - int(value*100)
	}
}
