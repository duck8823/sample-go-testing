package app

import (
	"net"
	"net/http"
	"testing"
	"io/ioutil"
	"net/url"
)

func Test_RoutingWithStartServer(t *testing.T) {
	addr := randomAddress(t)

	s := CreateServer()
	go func() {
		s.Start(addr.String())
	}()

	reqUrl := &url.URL{
		Scheme: "http",
		Host:   addr.String(),
		Path:	"hello",
		RawQuery: "name=duck",
	}
	resp, err := http.Get(reqUrl.String())
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	actual := string(body)
	expected := "Hello duck."
	if actual != expected {
		t.Errorf("got: %s\nwont: %s", actual, expected)
	}
}

func randomAddress(t *testing.T) net.Addr {
	t.Helper()

	listener, err := net.Listen("tcp", ":0")
	listener.Close()

	if err != nil {
		t.Fatal(err)
	}
	return listener.Addr()
}
