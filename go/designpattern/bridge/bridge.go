package bridge

import "fmt"

// bridge 桥接模式，多个维度组合

type ICoffee interface {
	OrderCoffee()
}

// LargeCoffee 大杯咖啡
type LargeCoffee struct {
	ICoffeeAddtion
}

type MediumCoffee struct {
	ICoffeeAddtion
}

func (lc *LargeCoffee) OrderCoffee() {
	fmt.Println("大杯咖啡")
	lc.AddSomething()
}

func (mc *MediumCoffee) OrderCoffee() {
	fmt.Println("中杯咖啡")
	mc.AddSomething()
}

type ICoffeeAddtion interface {
	AddSomething()
}

type Milk struct {
}

type Sugar struct {
}

func (milk Milk) AddSomething() {
	fmt.Println("加奶")
}

func (sugar Sugar) AddSomething() {
	fmt.Println("加糖")
}

type CoffeeAddtionType uint8

const (
	CoffeeAddMilk = iota
	CoffeeAddSugar
)

func NewCoffeeAddMilk() ICoffeeAddtion {
	return &Milk{}
}

func NewCoffeeAddSugar() ICoffeeAddtion {
	return &Sugar{}
}

var CoffeeAddtionFuncMap = map[CoffeeAddtionType] func() ICoffeeAddtion {
	CoffeeAddMilk:NewCoffeeAddMilk,
	CoffeeAddSugar: NewCoffeeAddSugar,
}

func NewCoffeeAddtion(addtionType CoffeeAddtionType) ICoffeeAddtion {
	if handler, ok := CoffeeAddtionFuncMap[addtionType]; ok {
		return handler()
	}
	return nil
}

type CoffeeType uint8
const (
	CoffeeCupLarge = iota
	CoffeeCupMedium
)

var CoffeeFuncMap = map[CoffeeType] func (addtionType ICoffeeAddtion) ICoffee {
	CoffeeCupLarge:NewLargeCoffee,
	CoffeeCupMedium: NewMediumCoffee,
}

func NewLargeCoffee(addtion ICoffeeAddtion) ICoffee  {
	return &LargeCoffee{addtion}
}

func NewMediumCoffee(addtion ICoffeeAddtion) ICoffee {
	return &MediumCoffee{addtion}
}


func NewCoffee(cupType CoffeeType, coffeeAddtion ICoffeeAddtion) ICoffee {
	if handler, ok := CoffeeFuncMap[cupType]; ok {
		return handler(coffeeAddtion)
	}
	return nil
}