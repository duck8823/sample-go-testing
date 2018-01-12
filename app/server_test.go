package app

import (
	"fmt"
	"github.com/duck8823/sample-go-testing/app/mock_server"
	"github.com/golang/mock/gomock"
	"net"
	"net/http"
	"testing"
)

func Test_RoutingWithStartServer(t *testing.T) {
	port := randomPort(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockController := mock_server.NewMockController(ctrl)
	mockController.EXPECT().Get(gomock.Any()).Times(1)

	s := Create(mockController)
	go func() {
		s.Start(port)
	}()

	_, err := http.Get(fmt.Sprintf("http://localhost%s/hello", port))
	if err != nil {
		t.Fatal(err)
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
