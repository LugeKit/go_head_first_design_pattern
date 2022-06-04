package main

import (
	"fmt"

	"github.com/lugekit/design_pattern/decorator"
	"github.com/lugekit/design_pattern/observer"
	"github.com/lugekit/design_pattern/singleton"
	"github.com/lugekit/design_pattern/strategy"
)

const dashLine = "-------------"

func main() {
	runOneMain(strategy.Main, "Strategy")
	runOneMain(observer.Main, "Observer")
	runOneMain(decorator.Main, "Decorator")
	runOneMain(singleton.Main, "Singleton")
}

func runOneMain(f func(), pattern string) {
	fmt.Printf("%s:\n", pattern)
	f()
	fmt.Println(dashLine)
}
