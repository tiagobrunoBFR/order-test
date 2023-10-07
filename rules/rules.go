package rules

import "order-test/entity"

type Rule interface {
	Apply(order *entity.Order)
}

type DiscountByPaymentMethod struct {
	MethodExpected     string
	DiscountPercentage float64
}

func (d *DiscountByPaymentMethod) Apply(o *entity.Order) {
	if d.MethodExpected == o.Payment.Method {
		var value float64
		value = float64(o.Payment.Value / 100)
		value = value * d.DiscountPercentage
		o.Payment.Value = o.Payment.Value - int(value*100)
	}
}

type LabelByCategory struct {
	CategoryExpected string
	Label            string
}

func (lc *LabelByCategory) Apply(o *entity.Order) {
	if o.Product.Category == lc.CategoryExpected {
		o.AddLabel(lc.Label)
	}
}

type LabelByValue struct {
	MaxValue int
	Label    string
}

func (lb *LabelByValue) Apply(o *entity.Order) {
	if o.Payment.Value > lb.MaxValue {
		o.AddLabel(lb.Label)
	}
}
