package animals

import "testing"

func TestDuck_name(t *testing.T) {
	duck := &Duck{"tarou"}
	actual := duck.name
	expected := "tarou"
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}
