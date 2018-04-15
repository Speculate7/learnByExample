package main

import "fmt"
import "math"

const circumferenceConst = 2 * math.Pi

func main() {
	/*
	// A `const` statement can appear anywhere a `var`
    // statement can.
    const n = 500000000

    // Constant expressions perform arithmetic with
    // arbitrary precision.
    const d = 3e20 / n
    fmt.Println(d)

    // A numeric constant has no type until it's given
    // one, such as by an explicit cast.
    fmt.Println(int64(d))

    // A number can be given a type by using it in a
    // context that requires one, such as a variable
    // assignment or function call. For example, here
    // `math.Sin` expects a `float64`.
    fmt.Println(math.Sin(n))
*/
    //Go has many constants already defined, for example `math.Pi`.
	//Print out the perimeter distance around a circle that has a radius of 5 feet.
	//The formula for perimeter distance is `2 * math.Pi * radius`
	radius := 5
	perimeter := circumferenceConst * float64(radius)

	//Also output the area of the same circle. The formula for area is math.Pi * math.Pow(radius,2)
	//Pow returns x**y, the base-x exponential of y. docs: func Pow(x, y float64) float64
	//area := math.Pi * math.Pow(radius,2)
	fmt.Printf("The circumference of a circle with a radius of 5ft is %.2fft\n",perimeter)
	//fmt.Printf("The area of a circle with a radius of 5ft is %.2fft^2\n",area)
	
}