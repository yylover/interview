package singleton

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	t.Run("singleton:", singletonTest)
}

func singletonTest(t *testing.T) {
	instance1 := GetInstance("ins1")
	fmt.Println(instance1.Title)
	instance2 := GetInstance("ins2")
	fmt.Println(instance2.Title)
}