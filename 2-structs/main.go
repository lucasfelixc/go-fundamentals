package main

import "fmt"

// Struct Order
type Order struct {
	ID       string
	Price    float64
	Quantity int
}

// Struct method to set the price
func (o *Order) setPrice(price float64) {
	o.Price = price
	fmt.Println("Price dentro do setPrice: ", o.Price)
}

// Struct method to get the total
func (o Order) getTotal() float64 {
	return o.Price * float64(o.Quantity)
}

func main() {
	order := Order{
		ID:       "123",
		Price:    10.0, // Here the price is defined with value 10
		Quantity: 5,
	}

	order.setPrice(20)                             // This method references the value of the original struct, and change the price value
	fmt.Println(order)                             // Here order.price is 20
	fmt.Println("Preço total: ", order.getTotal()) // This method multiplies price (which is 20) by quantity value (which is 5), so it will be 100
}
