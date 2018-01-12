package animals_test

import (
	"github.com/duck8823/sample-go-testing/animals"
	"github.com/duck8823/sample-go-testing/foods"
	"testing"
)

func TestDuck_05(t *testing.T) {
	t.Run("it says quack", func(t *testing.T) {
		duck := createInstance(t)

		actual := duck.Say()
		expected := "tarou says quack"
		if actual != expected {
			t.Errorf("got: %v\nwant: %v", actual, expected)
		}
	})

	t.Run("it ate apple", func(t *testing.T) {
		duck := createInstance(t)

		apple := foods.NewApple("sunfuji")

		actual := duck.Eat(apple)
		expected := "tarou ate sunfuji"
		if actual != expected {
			t.Errorf("got: %v\nwant: %v", actual, expected)
		}
	})
}

func createInstance(tb testing.TB) *animals.Duck {
	tb.Helper()

	duck := animals.NewDuck("tarou")

	return duck
}
