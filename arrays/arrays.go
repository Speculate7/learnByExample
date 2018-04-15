// In Go, an _array_ is a numbered sequence of elements of a
// specific length.

package main

import "fmt"

func main() {

	var a[10]int
	//"another way to creat an array" a := make([]int,10)
	fmt.Println("emp:", a)
	for i:=0;i<10;i++{
		a[i]= i+1
		fmt.Println("i =", a[i])
	}
	fmt.Println(a)

	p := make([]bool,10)
	fmt.Printf("%#v\n",p)

	

/*
	s := make([]int,0)
	fmt.Println("emp:", s)


	for i:=0;i<=10;i++{ 
		s = append(s, i+1)
		fmt.Println("apd:", s)
	}
	fmt.Println(s)
	fmt.Println(s[:6])
	fmt.Printf("%T\n",s[:6])//spits out the data type
	fmt.Printf("%#v\n",s[:6])//returns type and value
*/
}