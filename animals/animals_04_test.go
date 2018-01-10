package animals_test

import (
	"github.com/duck8823/sample-go-testing/animals"
	"testing"
)

func TestDuck_04(t *testing.T) {

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
			"it is named tarou",
			func(t *testing.T) {
				actual := duck.Name
				expected := "tarou"
				if actual != expected {
					t.Errorf("got: %v\nwant: %v", actual, expected)
				}
			},
		},
	} {
		// BeforeEach にあたる処理
		duck = &animals.Duck{"tarou"}

		// テストケースの実行
		t.Run(testcase.name, testcase.call)

		// AfterEach にあたる処理
	}
}
