package main

func run() {
	for i := 2; i <= 10; i++ {

		//if i != 5 {
		//	println(i)
		//}

		if i == 5 {
			continue
		}

		println(i)
	}
}

func loop2() {
	i := 10
	for i >= 5 {
		println(i)
		i--
	}
}

func loop3() {
	i := 1
	for i <= 5 {
		println(i)
		i++
	}
}

func loop4() {
	coffees := []string{"Cappuccino", "Latte", "Espresso"}

	for _, coffee := range coffees {
		println(coffee)
	}
}

func loop5() {
	coffees := []string{"Cappuccino", "Latte", "Espresso"}

	// for index, _ := range coffees {
	// 	println(index)
	// }

	for index := range coffees {
		println(index)
	}
}

func loop6() {
	coffees := []string{"Cappuccino", "Latte", "Espresso"}

	// for index, coffee := range coffees {
	// 	if index == 1 {
	// 		println(index, coffee)
	// 		break
	// 	}
	// }

	println(1, coffees[1])
}

func loop7() {
	outerArray := []int{1, 2}
	innerArray := []int{1, 2}
	for _, outvalue := range outerArray {
		for _, inValue := range innerArray {
			println(outvalue, inValue)
		}
	}
}

func loop8() {
	totalStarCount := 4
	star := 1

	for star <= totalStarCount {
		printStar := 1
		for printStar <= star {
			print("* ")
			printStar++
		}
		println()
		star++
	}
}

func loop9() {
	totalStarCount := 4

	for totalStarCount >= 0 {
		printStar := 1
		for printStar <= totalStarCount {
			print("* ")
			printStar++
		}
		println()
		totalStarCount--
	}
}

func loop10() {
	outerArray := []int{1, 2}
	innerArray := []int{1, 2}
	for _, outvalue := range outerArray {
		for _, inValue := range innerArray {
			if outvalue == 2 && inValue == 2 {
				break
			}
			println(outvalue, inValue)
		}
	}
}

func main() {
	//run()
	//loop2()
	//loop3()
	//loop4()
	//loop5()
	//loop6()
	//loop7()
	//loop8()
	//loop9()
	loop10()

}
