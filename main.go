package main

import "github.com/lugekit/design_pattern/strategy"

func main() {
	duck := strategy.NewNoobStrategyDuck()
	duck.PerformFly()
	duck.PerformSquak()
}
