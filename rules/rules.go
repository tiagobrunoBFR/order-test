package rules

import "order-test/entity"

type Rule struct {
	Order *entity.Order
}

func (r *Rule) DiscountByPaymentMethod(methodExpected string, discountPercentage float64) {
	if methodExpected == r.Order.Payment.Method {
		var value float64
		value = float64(r.Order.Payment.Value / 100)
		value = value * discountPercentage
		r.Order.Payment.Value = r.Order.Payment.Value - int(value*100)
	}
}

func (r *Rule) LabelByCategory(categoryExpected string, label string) {
	if r.Order.Product.Category == categoryExpected {
		r.Order.AddLabel(label)
	}
}

func (r *Rule) LabelByValue(maxValue int, label string) {
	if r.Order.Payment.Value > maxValue {
		r.Order.AddLabel(label)
	}
}
