package animals_test

import (
	"github.com/duck8823/sample-go-testing/animals"
	"github.com/duck8823/sample-go-testing/foods"
	"testing"
)

func TestDuck(t *testing.T) {
	duck := animals.NewDuck("tarou")

	t.Run("it says quack", func(t *testing.T) {
		actual := duck.Say()
		expected := "tarou says quack"
		if actual != expected {
			t.Errorf("got: %v\nwant: %v", actual, expected)
		}
	})

	t.Run("it ate apple", func(t *testing.T) {
		apple := foods.NewApple("sunfuji")

		actual := duck.Eat(apple)
		expected := "tarou ate sunfuji"
		if actual != expected {
			t.Errorf("got: %v\nwant: %v", actual, expected)
		}
	})
}
