package animals_test

import (
	"github.com/duck8823/sample-go-testing/animals"
	"github.com/duck8823/sample-go-testing/foods"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	println("before all...")

	code := m.Run()

	println("after all...")

	os.Exit(code)
}

func TestDuck_Say_03(t *testing.T) {
	duck := animals.NewDuck("tarou")
	actual := duck.Say()
	expected := "tarou says quack"
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}

func TestDuck_Eat(t *testing.T) {
	duck := animals.NewDuck("tarou")
	apple := foods.NewApple("sunfuji")

	actual := duck.Eat(apple)
	expected := "tarou ate sunfuji"
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}
