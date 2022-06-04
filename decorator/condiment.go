package decorator

type CondimentDecorator struct {
	beverage Beverage
}

func (d *CondimentDecorator) Description() string {
	return d.beverage.Description()
}

func (d *CondimentDecorator) Cost() float64 {
	return d.beverage.Cost()
}

type Milk struct {
	CondimentDecorator
}

func AddMilk(beverage Beverage) Beverage {
	return &Milk{
		CondimentDecorator{
			beverage: beverage,
		},
	}
}

func (m *Milk) Description() string {
	return "Milk " + m.CondimentDecorator.Description()
}

func (m *Milk) Cost() float64 {
	return 0.20 + m.CondimentDecorator.Cost()
}
