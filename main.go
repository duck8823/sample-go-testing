package main

import (
	"fmt"
	"github.com/duck8823/sample-go-testing/animals"
	"github.com/duck8823/sample-go-testing/app"
	"github.com/duck8823/sample-go-testing/foods"
	"os"
)

func main() {
	optC := app.ParseFlag(os.Args[1:]...)
	fmt.Printf("option c: %s\n", optC)

	duck := &animals.Duck{"tarou"}
	apple := foods.NewApple("sunfuji")
	println(duck.Eat(apple))

	server := app.CreateServer()
	server.Start(":8080")
}
