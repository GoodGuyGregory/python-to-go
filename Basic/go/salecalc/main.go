package main

import "fmt"

var orignalprice float64
var discount float64

func main() {

	fmt.Println("enter the price of the item")
	_, err := fmt.Scan(&originalprice)

	fmt.Println("enter the precentage off")
	_, err = fmt.Scan(&discount)

	if err == nil {
		fmt.Println("the price is", originalprice)
		// var precentDiscount int
	} else {
		fmt.Println(err)
	}

	if err == nil {
		// Convert to precentageOff
		var precentageoff = discount / 100
		// Calculate adjustment
		var adjustedprice = originalprice * precentageoff
		var saleprice = orignalprice - adjustedprice
		fmt.Println("the precentage off is", saleprice)
	} else {
		fmt.Println(err)
	}

}
