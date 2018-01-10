package animals_test

import "testing"
import "github.com/duck8823/sample-go-testing/animals"

func TestDuck_Say(t *testing.T) {
	duck := &animals.Duck{"tarou"}
	actual := duck.Say()
	expected := "tarou says quack"
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}
