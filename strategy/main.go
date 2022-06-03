package strategy

func Main() {
	duck := NewStrategyNoobDuck()
	duck.PerformSquak()
	duck.PerformFly()
}
