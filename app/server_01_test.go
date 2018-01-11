package app

import (
	"github.com/duck8823/sample-go-testing/app/mock_server"
	"github.com/golang/mock/gomock"
	"net/http/httptest"
	"testing"
)

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
