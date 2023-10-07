package rules

import "order-test/entity"

func DiscountByPaymentMethod(methodExpected string, discountPercentage float64, o *entity.Order) {
	if methodExpected == o.Payment.Method {
		var value float64
		value = float64(o.Payment.Value / 100)
		value = value * discountPercentage
		o.Payment.Value = o.Payment.Value - int(value*100)
	}
}
