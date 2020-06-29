package main

import (

"fmt"

)


func main()  {

	var num_items int = 100

	fmt.Println(num_items)

	
//var price1 float64 = 9.99 
//var price2 float64 = 5.50 

// how to declare 2 variables at the same time:        
//   var price1, price2 float64 = 9.99, 5.50

// how to decalre 2 variables and assing them values at the same time      
price1, price2 := 9.99, 5.50 

fmt.Println(price1, price2)



}