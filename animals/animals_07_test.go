package animals_test

import (
	"github.com/duck8823/sample-go-testing/animals"
	"github.com/duck8823/sample-go-testing/foods/mock_foods"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestDuck_Eat_02(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	food := mock_foods.NewMockFood(ctrl)
	food.EXPECT().Name().Return("kougyoku")

	duck := animals.NewDuck("tarou")
	actual := duck.Eat(food)
	expected := "tarou ate kougyoku"
	if actual != expected {
		t.Errorf("got: %s\nwont: %s", actual, expected)
	}
}
