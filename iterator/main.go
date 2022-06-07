package iterator

import "fmt"

func Main() {
	menu := SliceMenu{
		items: []int{1, 2, 3, 4},
	}
	it := menu.Iterator()

	for it.HasNext() {
		fmt.Printf("%d ", it.Next())
	}
	fmt.Println()
}
