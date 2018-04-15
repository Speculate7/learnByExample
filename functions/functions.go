
package main

import "fmt"

func loopForDays()string{
	 for i := 0; i <= 10; i++{
	 	switch {
	 		case i%2 == 0:
	 			resp := fmt.Printf("%d is even\n", i)
	 		default:
	 			resp := fmt.Printf("%d is odd\n", i)
	 	}
	 	return resp
	 }
}

func main(){
	loopForDays(10)
	
}