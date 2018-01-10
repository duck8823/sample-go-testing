package animals_test

import (
	"github.com/duck8823/sample-go-testing/animals"
	"testing"
	"time"
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

	t.Run("it is named tarou", func(t *testing.T) {
		duck := createInstance(t)

		actual := duck.Name
		expected := "tarou"
		if actual != expected {
			t.Errorf("got: %v\nwant: %v", actual, expected)
		}
	})
}

func createInstance(tb testing.TB) *animals.Duck {
	tb.Helper()

	duck := &animals.Duck{"tarou"}
	time.Sleep(5 * time.Second)
	tb.Error("前処理で失敗しました.")

	return duck
}
