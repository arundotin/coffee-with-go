package main

func a1() {
	var coffees [3]string

	coffees[0] = "Cappuccino"
	coffees[1] = "Latte"
	coffees[2] = "Espresso"

	println(coffees[1])
	println(len(coffees))
}
func a2() {
	coffees := []string{"Cappuccino", "Latte", "Espresso"}

	println(coffees[1])
	println(len(coffees))
}

func a3() {
	coffees := []string{"Cappuccino", "Latte"}

	coffees = append(coffees, "Espresso")

	coffees[2] = "Filter Coffee"

	println(len(coffees))
	println(coffees[2])
}

func a4() {
	coffees := make([]string, 2, 5)

	//println(len(coffees))
	//println(cap(coffees))
	println(coffees[0])
	println(coffees[1])

	coffees[1] = "Arun"

	println(coffees[0])
	println(coffees[1])
	println(len(coffees))
	coffees = append(coffees, "Kumar")

	println(len(coffees))
}

func a5() {
	coffees := []string{"Cappuccino", "Latte", "Espresso", "Mocha"}

	selected := coffees[1:3]

	println(len(selected))
	println(selected[0])
	println(selected[1])

}

func a6() {
	coffees := []string{"Cappuccino", "Latte", "Espresso", "Mocha"}

	selected := coffees[1:3]

	println(coffees[1])

	selected[0] = "Filter Coffee"

	println(coffees[1])
	println(selected[0])
	println(selected[1])

	// slices shares underlying array
}

func a7() {
	coffees := []string{"Cappuccino", "Latte", "Espresso", "Mocha"}
	selected := make([]string, 2)
	copy(selected, coffees[1:3])

	println(selected[0])
	println(selected[1])
	selected[0] = "Filter Coffee"
	println(coffees[1]) // coffees wont change, because it doesnt share the underlying array
	println(selected[0])

}

func a8() {
	coffees := []string{"Cappuccino", "Latte", "Espresso", "Mocha"}
	selected := make([]string, 2)
	copy(selected, coffees[1:3])

	println(selected[0])
	println(selected[1])
	selected[0] = "Filter Coffee"
	println(coffees[1]) // coffees wont change, because it doesnt share the underlying array
	println(selected[0])

	animals := []string{"Dog", "Cow"}

	copy(selected, animals)

	println(selected[0])

}

func a9() {
	coffees := []string{"Cappuccino", "Latte", "Espresso"}

	i := 0

	for i < len(coffees) {
		println(coffees[i])
		i++
	}
}

func a10() {
	coffees := []string{"Cappuccino", "Latte", "Espresso"}

	for _, coffee := range coffees {
		println(coffee)
	}
}

func main() {
	//a1()
	//a2()
	//a3()
	//a4()
	//a5()
	//a6()
	//a7()
	//a8()
	//a9()
	a10()
}
