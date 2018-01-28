package animals_test

import (
	"github.com/duck8823/sample-go-testing/animals"
	"github.com/duck8823/sample-go-testing/foods"
	"testing"
)

func BenchmarkDuck_Say(b *testing.B) {
	duck := animals.NewDuck("tarou")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		duck.Say()
	}
	b.StopTimer()
}

func BenchmarkDuck_Eat(b *testing.B) {
	duck := animals.NewDuck("tarou")
	food := foods.NewApple("sunfuji")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		duck.Eat(food)
	}
	b.StopTimer()
}
