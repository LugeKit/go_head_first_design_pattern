package decorator

type Beverage interface {
	Description() string
	Cost() float64
}

type Expresso struct{}

func (e *Expresso) Description() string {
	return "Expresso"
}

func (e *Expresso) Cost() float64 {
	return 10.0
}
