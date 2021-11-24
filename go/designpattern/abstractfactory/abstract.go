package abstractfactory

import "fmt"


/**
抽象工厂，类似实现了多态的特征
 */

type FruitFactory interface {
	CreateFruit() Fruit
}

type Fruit interface {
	Eat()
}

type AppleFactory struct {
}

func (af *AppleFactory) CreateFruit() Fruit {
	return &Apple{}
}

type Apple struct {
}

func (ap *Apple) Eat() {
	fmt.Println("apple eat")
}


type BananaFactory struct {
}

func (af *BananaFactory) CreateFruit() Fruit {
	return &Banana{}
}

type Banana struct {
}

func (ba *Banana) Eat() {
	fmt.Println("Banana eat")
}

func eat(ff FruitFactory) {
	fruit := ff.CreateFruit()
	fruit.Eat()
}

//
// func main() {
// 	appleFactory := &AppleFactory{}
// 	eat(appleFactory)
//
// 	bananaFactory := &BananaFactory{}
// 	eat(bananaFactory)
// }