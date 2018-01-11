package foods_test

import (
	"testing"
	"github.com/duck8823/sample-go-testing/foods"
)

func TestNewApple_Name(t *testing.T) {
	apple := foods.NewApple("sunfuji")
	actual := apple.Name()
	expected := "sunfuji"
	if actual != expected {
		t.Errorf("got: %s\nwont%s", actual, expected)
	}
}