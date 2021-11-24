package simplefactory

import "testing"

func Test(t *testing.T) {
	t.Run("simple_factory:", ProduceFruitAndEat)
}

func ProduceFruitAndEat(t *testing.T) {
	var apple, banana Fruit
	apple = ProduceFruit(FruitTypeApple)
	banana = ProduceFruit(FruitTypeBanana)

	apple.Eat()
	banana.Eat()
}