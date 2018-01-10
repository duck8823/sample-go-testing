package animals_test

import "testing"
import "github.com/duck8823/sample-go-testing/animals"

func TestDuck(t *testing.T) {
	duck := &animals.Duck{"tarou"}

	t.Run("it says quack", func(t *testing.T) {
		actual := duck.Say()
		expected := "tarou says quack"
		if actual != expected {
			t.Errorf("got: %v\nwant: %v", actual, expected)
		}
	})

	t.Run("it is named tarou", func(t *testing.T) {
		actual := duck.Name
		expected := "tarou"
		if actual != expected {
			t.Errorf("got: %v\nwant: %v", actual, expected)
		}
	})
}
