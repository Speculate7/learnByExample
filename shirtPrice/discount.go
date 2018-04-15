package main 

import "fmt"

var price = 10.00


func main() {
	//price := 10.00 - must be defined inside the scope if using this method
	discount30 := 0.30
	discount25 := 0.25
	discount10 := 0.10
	
	salesPrice10 := price - (price * discount10)
	salesPrice25 := price - (price * discount25)
	salesPrice30 := price - (price * discount30)
	
	fmt.Printf("The under armour shirt costs $%.2f with a 10%% discount\n",salesPrice10)
	fmt.Printf("The under armour shirt costs $%.2f with a 25%% discount\n",salesPrice25) 
	fmt.Printf("The under armour shirt costs $%.2f with a 30%% discount\n",salesPrice30) 
}