package main

import "fmt"

func main() {

	originalprice := 0.0
	discount := 0.0

	fmt.Println("enter the price of the item")
	// Scanf allows for flloat reading
	_, err := fmt.Scanf("%f", &originalprice)

	if err != nil {
		fmt.Printf("Invalid price: $s\n", err)
		// return will stop the program which is convention in go
		return
	}

	fmt.Println("enter the precentage off")
	_, err = fmt.Scanf("%f", &discount)

	if err != nil {
		fmt.Printf("Invalid discount: $s\n", err)
		// var precentDiscount int
		return
	}
	// declares and creates the variable
	percentOff := discount / 100.0
	adjustedPrice := originalprice * percentOff
	salePrice := originalprice - adjustedPrice

	// %.2f formats the decimal points for the values
	fmt.Printf("The percent off is %.2f.\n The Price will be: %.2f\n", percentOff, salePrice)
}
