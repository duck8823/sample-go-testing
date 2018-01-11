package animals_test

import "testing"
import (
	"github.com/duck8823/sample-go-testing/animals"
	"os"
)

func TestMain(m *testing.M) {
	println("before all...")

	code := m.Run()

	println("after all...")

	os.Exit(code)
}

func TestDuck_Say_03(t *testing.T) {
	duck := &animals.Duck{"tarou"}
	actual := duck.Say()
	expected := "tarou says quack"
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}
