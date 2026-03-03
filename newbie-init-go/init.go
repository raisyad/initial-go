package main

import (
	"fmt"
	"time"
	"math"
	"math/rand"
)

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

	b, c, d, _, f := 0,1,2,3, "adadeh"
	fmt.Println(b,c,d,f)

	//For hexadecimalm, binary and default int
	x,y,z := 0,3,5
	fmt.Printf("%v \t %b \t %X\n",x,x,x);
	fmt.Printf("%v \t %b \t %X\n",y,y,y);
	fmt.Printf("%v \t %b \t %X\n",z,z,z);

	//For time
	fmt.Println(time.Now())

	//For random number
	fmt.Println(rand.Intn(10))

	//For math sqrt
	fmt.Println(math.Sqrt(9))

	//For pi
	fmt.Println(math.Pi)
}
