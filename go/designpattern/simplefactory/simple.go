package simplefactory

import "fmt"

type Fruit interface {
	Eat()
}

type Apple struct {
}

func (ap *Apple) Eat() {
	fmt.Println("Apple eat")
}

type Banana struct {
}

func (ba *Banana) Eat() {
	fmt.Println("Banana eat")
}

type FruitType uint8

const (
	FruitTypeApple = iota
	FruitTypeBanana
)

var FruitFuncMap = map[FruitType] func() Fruit {
	FruitTypeApple : produceApple,
	FruitTypeBanana: produceBanana,
}

//ProduceFruit 抽象工厂
func ProduceFruit(fruitType FruitType) Fruit {
	if handler, ok := FruitFuncMap[fruitType]; ok {
		return handler()
	}
	return nil
}

func produceApple() Fruit {
	return &Apple{}
}

func produceBanana() Fruit {
	return &Banana{}
}