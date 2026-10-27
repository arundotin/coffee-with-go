package main

import "fmt"

func makeCoffee(coffeeType string) error {
	fmt.Println("Preparing coffee of the type " + coffeeType)
	coffeePrice, err := getCoffeePrice(coffeeType)
	if err != nil {
		return err
	}

	fmt.Println("Price is", coffeePrice)
	return nil
}

func getCoffeeName() string {
	return "Cappuccino"
}

func getCoffeePrice(coffeeType string) (float64, error) {
	coffeePrices := map[string]float64{
		"Cappuccino": 150.0,
		"Latte":      120.0,
	}

	price, valuePresent := coffeePrices[coffeeType]

	if !valuePresent {
		return 0, fmt.Errorf("unknown coffee: %s", coffeeType)
	}

	return price, nil
}

func run() error {
	coffeeName := getCoffeeName()

	if err := makeCoffee(coffeeName); err != nil {
		return err
	}

	if err := makeCoffee("Latte"); err != nil {
		return err
	}

	if err := makeCoffee("Espresso"); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println("Application failed:", err)
	}
}
