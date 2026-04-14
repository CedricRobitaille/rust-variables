package main

import "fmt"

func stringInterpolation() {
	const name = "Cedric"
	const openRate = 5.2

	/*
		For string interpolation:
		- %v any value in default format ("i am %v years old", 23 OR "too many")
		- %s string values ("i am %s years old", "too many")
		- %d integer values ("i am %d years old", 23)
		- %f float values ("i am %.2f years old", 23.756)
			- %.2f denotes decimal points.
	*/
	msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent\n", name, openRate)

	fmt.Print(msg)
}

func main() {
	stringInterpolation()
}
