package foods

type Food interface {
	Name() string
}

type Apple struct {
	cultivar string
}

func NewApple(cultivar string) *Apple {
	return &Apple{cultivar}
}

func (apple *Apple) Name() string {
	return apple.cultivar
}
