package animals

import "testing"

func TestDuck_Say(t *testing.T) {
	duck := &Duck{"tarou"}
	actual := duck.Say()
	expected := "tarou says quack"
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}
