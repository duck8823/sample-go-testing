package app

import (
	"net/http/httptest"
	"testing"
)

func Test_RoutingWitHttpTestSimulate(t *testing.T) {
	s := CreateServer()

	req := httptest.NewRequest("GET", "/hello?name=duck", nil)
	rec := httptest.NewRecorder()

	s.ServeHTTP(rec, req)

	actual := rec.Body.String()
	expected := "Hello duck."
	if actual != expected {
		t.Errorf("got: %s\nwont: %s", actual, expected)
	}
}
