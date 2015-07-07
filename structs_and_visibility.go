package main

import "fmt"

type Beer struct {
	Brewery  string
	Name     string
	Style    string
	calories int
}

func (b Beer) ToString() string {
	return fmt.Sprintf("%v %v", b.Brewery, b.Name)
}

func (b Beer) GetCalories() int {
	return b.superSecretFunction()
}

func (b Beer) superSecretFunction() int {
	return b.calories
}

func main() {
	barelyBeer := Beer{"Michelob", "Ultra", "Light", 60}

	fmt.Println(barelyBeer.ToString(), "has",
		barelyBeer.GetCalories(), "calories.")

	// this is available because we're in the same package
	// but if we had imported from another package then it
	// wouldn't be
	fmt.Println(barelyBeer.superSecretFunction())
}
