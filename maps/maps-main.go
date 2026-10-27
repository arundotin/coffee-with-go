package main

import "sort"

func m1() {
	coffees := map[string]int{
		"Cappuccino": 120,
		"Latte":      150,
		"Espresso":   100,
	}

	println(coffees["Cappuccino"])

	// Add a new entry
	coffees["Mocha"] = 180

	// Update an existing entry
	coffees["Latte"] = 160

	println(coffees["Mocha"])
	println(coffees["Latte"])
}

func m2() {
	type Employee struct {
		ID   int
		Name string
	}

	employees := map[Employee]string{}

	employees[Employee{ID: 101, Name: "Arun"}] = "Engineering"
	employees[Employee{ID: 101, Name: "Arun"}] = "Management"

	println(len(employees)) // 1
}

func m3() {
	coffees := map[string]int{
		"Espresso":   100,
		"Cappuccino": 120,
		"Latte":      150,
	}

	//create a slice
	coffeeNames := []string{}

	// Collect keys into a slice
	for coffeeName := range coffees {
		println(coffeeName)
		coffeeNames = append(coffeeNames, coffeeName)
	}

	sort.Strings(coffeeNames)

	for index, coffeeName := range coffeeNames {
		println(index, coffeeName)
	}

}

func m4() {
	coffees := map[string]int{
		"Cappuccino": 120,
		"Latte":      150,
	}

	delete(coffees, "Cappuccino")

	println(len(coffees))
}

func m7() {
	var coffees map[string]int

	println(coffees == nil)
	println(len(coffees))

	coffees = make(map[string]int)
	coffees["Espresso"] = 100

	println(coffees == nil)
	println(len(coffees))
}

func main() {
	//m1()
	//m2()
	//m3()
	//m4()
	m7()
}
