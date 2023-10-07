package rules

import "order-test/entity"

func LabelByValue(maxValue int, label string, o *entity.Order) {
	if o.Payment.Value > maxValue {
		o.AddLabel(label)
	}
}
