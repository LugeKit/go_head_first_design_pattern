package template

import "fmt"

type Preparation interface {
	BoilWater()
	Brew()
	PourInCup()
	AddCondiments()
}

type BasePreparation struct{}

func (p *BasePreparation) BoilWater() {
	fmt.Println("boil water")
}

func (p *BasePreparation) Brew() {
	// do nothing
}

func (p *BasePreparation) PourInCup() {
	fmt.Println("pour in cup")
}

func (p *BasePreparation) AddCondiments() {
	// do nothing
}

type Coffee struct {
	*BasePreparation
}

func (c *Coffee) Brew() {
	fmt.Println("coffee brew")
}

func (c *Coffee) AddCondiments() {
	fmt.Println("coffee add condiments")
}

type Tea struct {
	*BasePreparation
}

func (t *Tea) Brew() {
	fmt.Println("tea brew")
}

func (t *Tea) AddCondiments() {
	fmt.Println("tea add condiments")
}

type CaffeineBeverage struct {
	p Preparation
}

func (cb *CaffeineBeverage) PrepareRecipe() {
	cb.p.BoilWater()
	cb.p.Brew()
	cb.p.PourInCup()
	cb.p.AddCondiments()
}

func NewCoffee() *CaffeineBeverage {
	return &CaffeineBeverage{p: &Coffee{}}
}

func NewTea() *CaffeineBeverage {
	return &CaffeineBeverage{p: &Tea{}}
}
