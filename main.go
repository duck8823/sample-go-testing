package main

import (
	"github.com/duck8823/sample-go-testing/animals"
	"github.com/duck8823/sample-go-testing/foods"
)

func main() {
	duck := &animals.Duck{}
	apple := &foods.Apple{}
	println(duck.Eat(apple))
}
