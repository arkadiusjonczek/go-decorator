package main

import "fmt"

type IPizza interface {
    getPrice() int
}

var _ IPizza = (*VeggieMania)(nil)

type VeggieMania struct {
}

func (p *VeggieMania) getPrice() int {
    return 15
}

var _ IPizza = (*TomatoTopping)(nil)

type TomatoTopping struct {
    pizza IPizza
}

func (c *TomatoTopping) getPrice() int {
    pizzaPrice := c.pizza.getPrice()
    return pizzaPrice + 7
}

func main() {
	var pizza1 IPizza
	pizza1 = &VeggieMania{}
	fmt.Println(pizza1.getPrice())
	var pizza2 IPizza
	pizza2 = &TomatoTopping{pizza1}
	fmt.Println(pizza2.getPrice())
}
