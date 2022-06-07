package template

func Main() {
	coffee := NewCoffee()
	coffee.PrepareRecipe()

	tea := NewTea()
	tea.PrepareRecipe()
}
