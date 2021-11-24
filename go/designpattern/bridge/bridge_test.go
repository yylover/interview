package bridge

import (
	"testing"
)


func Test(t *testing.T) {
	t.Run("large-milk:", OrderLargeMilkCoffee)
}

func OrderLargeMilkCoffee(t *testing.T) {
	var coffeeAddtion ICoffeeAddtion
	coffeeAddtion = NewCoffeeAddtion(CoffeeAddMilk)
	var coffee ICoffee
	coffee = NewCoffee(CoffeeCupLarge, coffeeAddtion)
	coffee.OrderCoffee()
}