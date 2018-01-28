package animals

import (
	"fmt"
	"github.com/duck8823/sample-go-testing/foods"
)

type Duck struct {
	name string
}

func NewDuck(name string) *Duck {
	return &Duck{name}
}

func (duck *Duck) Say() string {
	return fmt.Sprintf("%s says quack", duck.name)
}

func (duck *Duck) Eat(food foods.Food) string {
	return fmt.Sprintf("%s ate %s", duck.name, food.Name())
}
