package server

import (
	"testing"
	"net/http"
	"github.com/duck8823/sample-go-testing/server/mock_server"
	"github.com/golang/mock/gomock"
)

func Test_Hello(t *testing.T) {
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
