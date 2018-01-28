package animals_test

import (
	"github.com/duck8823/sample-go-testing/animals"
	"github.com/duck8823/sample-go-testing/foods"
	"testing"
)

func TestDuck_04(t *testing.T) {

	println("before all...")
	var duck *animals.Duck

	for _, testcase := range []struct {
		name string
		call func(t *testing.T)
	}{
		{
			"it says quack",
			func(t *testing.T) {
				actual := duck.Say()
				expected := "tarou says quack"
				if actual != expected {
					t.Errorf("got: %v\nwant: %v", actual, expected)
				}
			},
		}, {
			"it ate apple",
			func(t *testing.T) {
				apple := foods.NewApple("sunfuji")

				actual := duck.Eat(apple)
				expected := "tarou ate sunfuji"
				if actual != expected {
					t.Errorf("got: %v\nwant: %v", actual, expected)
				}
			},
		},
	} {
		println("before each...")
		duck = animals.NewDuck("tarou")

		// テストケースの実行
		t.Run(testcase.name, testcase.call)

		println("after each...")
	}
	println("after all...")
}
