package main

import (
	"fmt"

	"github.com/lugekit/design_pattern/decorator"
	"github.com/lugekit/design_pattern/observer"
	"github.com/lugekit/design_pattern/strategy"
)

const dashLine = "-------------"

func main() {
	runOneMain(strategy.Main)
	runOneMain(observer.Main)
	runOneMain(decorator.Main)
}

func runOneMain(f func()) {
	f()
	fmt.Println(dashLine)
}
