package decorator

import "fmt"

func Main() {
	// get Expresso with Milk
	expresso := &Expresso{}
	expressoWithMilk := AddMilk(expresso)
	fmt.Printf("Description: %s\n", expressoWithMilk.Description())
	fmt.Printf("Cost: %v\n", expressoWithMilk.Cost())
}
