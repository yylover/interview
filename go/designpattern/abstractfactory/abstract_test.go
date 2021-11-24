package abstractfactory


import "testing"

func Test(t *testing.T) {
	t.Run("abstractfactory:", ProduceFruitAndEat)
}

func ProduceFruitAndEat(t *testing.T) {
	appleFactory := &AppleFactory{}
	eat(appleFactory)

	bananaFactory := &BananaFactory{}
	eat(bananaFactory)
}