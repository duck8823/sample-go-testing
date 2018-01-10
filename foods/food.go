package foods

type Food interface {
	Name() string
}

type Apple struct {
	name string
}

func NewApple(name string) *Apple {
	return &Apple{name}
}

func (apple *Apple) Name() string {
	return apple.name
}
