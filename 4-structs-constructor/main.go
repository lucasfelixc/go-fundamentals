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

func NewOrder() *Order {
	return &Order{}
}

func main() {
	order := NewOrder()
	order.ID = "123"
	order.Quantity = 5
	order.Price = 10

	order2 := order
	order2.Price = 50
	fmt.Println(order.Price)
}
