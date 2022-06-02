package strategy

import "fmt"

type duck interface {
	Squak()
	Fly()
}

type noobDuck struct {
}

func (d *noobDuck) Squak() {
	fmt.Println("I can't squak!")
}

func (d *noobDuck) Fly() {
	fmt.Println("I can't fly!")
}
