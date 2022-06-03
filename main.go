package main

import (
	"fmt"

	"github.com/lugekit/design_pattern/observer"
	"github.com/lugekit/design_pattern/strategy"
)

const dashLine = "-------------"

func main() {
	strategy.Main()
	fmt.Println(dashLine)
	observer.Main()
}
