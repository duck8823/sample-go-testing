package app

import (
	"github.com/duck8823/sample-go-testing/app/mock_server"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_RoutingWithStartServer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockController := mock_server.NewMockController(ctrl)
	mockController.EXPECT().Get(gomock.Any()).Times(1)

	s := Create(mockController)
	go func() {
		s.Start(":1234")
	}()

	_, err := http.Get("http://localhost:1234/hello")
	if err != nil {
		t.Fatal(err)
	}
}

func Test_RoutingWitHttpTest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockController := mock_server.NewMockController(ctrl)
	mockController.EXPECT().Get(gomock.Any()).Times(1)

	s := Create(mockController)
	server := httptest.NewServer(s)
	client := server.Client()

	_, err := client.Get(server.URL + "/hello")
	if err != nil {
		t.Fatal(err)
	}
}
