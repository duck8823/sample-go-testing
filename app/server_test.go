package app

import (
	"fmt"
	"net"
	"net/http"
	"testing"
	"io/ioutil"
)

func Test_RoutingWithStartServer(t *testing.T) {
	port := randomPort(t)

	s := CreateServer()
	go func() {
		s.Start(port)
	}()

	resp, err := http.Get(fmt.Sprintf("http://localhost%s/hello?name=duck", port))
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

func randomPort(t *testing.T) string {
	t.Helper()

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	addr := listener.Addr().String()
	_, port, err := net.SplitHostPort(addr)
	listener.Close()

	return fmt.Sprintf(":%s", port)
}
