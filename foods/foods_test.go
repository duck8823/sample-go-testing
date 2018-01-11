package foods_test

import (
	"github.com/duck8823/sample-go-testing/foods"
	"testing"
)

func TestNewApple_Name(t *testing.T) {
	apple := foods.NewApple("sunfuji")
	actual := apple.Name()
	expected := "sunfuji"
	if actual != expected {
		t.Errorf("got: %s\nwont%s", actual, expected)
	}
}
