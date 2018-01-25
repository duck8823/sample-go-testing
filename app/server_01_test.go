package app

import (
	"net/http/httptest"
	"testing"
	"io/ioutil"
)

func Test_RoutingWitHttpTest(t *testing.T) {
	s := CreateServer()
	server := httptest.NewServer(s)
	defer server.Close()

	client := server.Client()

	resp, err := client.Get(server.URL + "/hello?name=duck")
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	actual := string(body)
	expected := "Hello duck."
	if string(actual) != expected {
		t.Errorf("got: %s\nwont: %s", actual, expected)
	}
}
