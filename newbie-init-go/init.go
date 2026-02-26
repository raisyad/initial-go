package main

import "fmt"

func main() {
	const age, name = 22, "raisyad"
	var height int = 170
	var weight = 60
	var isMarried = false
	var pi float32 = 3.14
	var named string = "raisyad"

	fmt.Println("Hello World")

	fmt.Println(`Lorem Ipsum 
	dolor sit amet consectetur adipiscing elit`)

	fmt.Println("Age & Name : ", age, name)
	fmt.Println("Age :", age, ", Name :", name)
	fmt.Printf("Height : %d\n", height)
	fmt.Printf("Weight : %d\n", weight)
	fmt.Printf("Is Married : %t\n", isMarried)
	fmt.Printf("Pi : %.2f\n", pi)
	fmt.Printf("Name : %s\n", named)
}
