package strategy

import "fmt"

type squakBehavior interface {
	Squak()
}

type flyBehavior interface {
	Fly()
}

type strategyDuck interface {
	PerformSquak()
	PerformFly()
}

type baseStrategyDuck struct {
	squakExcutor squakBehavior
	flyExcutor   flyBehavior
}

func (d *baseStrategyDuck) PerformSquak() {
	d.squakExcutor.Squak()
}

func (d *baseStrategyDuck) PerformFly() {
	d.flyExcutor.Fly()
}

type strategyNoobDuck struct {
	baseStrategyDuck
}

func NewStrategyNoobDuck() strategyDuck {
	return &strategyNoobDuck{
		baseStrategyDuck: baseStrategyDuck{
			squakExcutor: &noSquakBehavior{},
			flyExcutor:   &noFlyBehavior{},
		},
	}
}

type noSquakBehavior struct {
}

func (sb *noSquakBehavior) Squak() {
	fmt.Println("I can't squak")
}

type noFlyBehavior struct{}

func (fb *noFlyBehavior) Fly() {
	fmt.Println("I can't fly")
}
