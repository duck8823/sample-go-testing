package app_test

import (
	"github.com/duck8823/sample-go-testing/app"
	"testing"
)

func TestParseFlag(t *testing.T) {
	t.Run("with argument, returns value", func(t *testing.T) {
		actual := app.ParseFlag("-c", "hello")
		expected := "hello"

		if actual != expected {
			t.Errorf("got: %s\nwont: %s", actual, expected)
		}
	})

	t.Run("with no argument, returns default value", func(t *testing.T) {
		actual := app.ParseFlag()
		expected := "default value"

		if actual != expected {
			t.Errorf("got: %s\nwont: %s", actual, expected)
		}
	})
}
