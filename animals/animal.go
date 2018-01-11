package animals

import (
	"fmt"
	"github.com/duck8823/sample-go-testing/foods"
)

type Animal interface {
	Say() string
	Eat(food foods.Food) string
}

type Duck struct {
	Name string
}

func (duck *Duck) Say() string {
	return fmt.Sprintf("%s says quack", duck.Name)
}

func (duck *Duck) Eat(food foods.Food) string {
	return fmt.Sprintf("%s ate %s", duck.Name, food.Name())
}
