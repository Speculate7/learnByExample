package main
import "fmt"

 func main() {
 //basic if/else statements. else if can also be used. If using boolean and a statement evaluates
 //to true the "else if" will not be executed.	
 /*
 	count := 10
 	for i := 0; i <= count; i++{
 		if i%2 == 0 {
 			fmt.Printf("%d is even\n",i)
 		} else {
 			fmt.Printf("%d is odd\n",i)
 		}
 	}
	*/
 //usinfg "switch"
 	for i := 0; i <= 10; i++{//condtions must be written in this order
		switch {//no need to specify "i" if used in condition
			case i < 5:
				fmt.Printf("%d is less than 5\n", i)
			case i == 5:
				fmt.Printf("%d is equal\n", i)
			default://no conditions for default
				fmt.Printf("%d is greater than 5\n", i)
		}
	}
/*
switch i {
	case 1,2,3,4:
		fmt.Println(i, "is less than 5")
	case 5:
		fmt.Println(i, "is equal to 5")
	case 6,7,8,9,10:
		fmt.Println(i, "is greater than 5")
}
*/
/*
switch i {
	case 1,2,3,4:
		fmt.Println(i, "is less than 5")
		continue
	case 5:
		fmt.Println(i, "is equal to 5")
		continue
}
	fmt.Println(i, "is greater than 5")
*/

}
