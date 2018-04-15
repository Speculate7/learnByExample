// `for` is Go's only looping construct. Here are
// three basic types of `for` loops.
//Write a loop with no condition that exits when the count is equal to 100.
package main

import "fmt"

func main() {
  n := 0
  for {
  	if n == 100 {
      break
    }
   n++ 
   fmt.Println(n)
  }
  fmt.Printf("I made it to %d!\n",n)

//utilizing the "return" also break the entire loop
  var j int
  for {
	fmt.Println(j)
	if j == 100 {
		return
	}
	j++
  }
}

/*
Other examples
// The most basic type, with a single condition.
    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }

    // A classic initial/condition/after `for` loop.
    for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }

    // `for` without a condition will loop repeatedly
    // until you `break` out of the loop or `return` from
    // the enclosing function.
    for {
        fmt.Println("loop")
        break
    }

    // You can also `continue` to the next iteration of
    // the loop.
    for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
*/