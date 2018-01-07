package animals

import "github.com/duck8823/sample-go-testing/foods"

type Animal interface {
	Say() string
	Eat(food foods.Food) string
}

type Duck struct {
}

func (duck *Duck) Say() string {
	return "hoge"
}

func (duck *Duck) Eat(food foods.Food) string {
	return "duck ate " + food.Name()
}
