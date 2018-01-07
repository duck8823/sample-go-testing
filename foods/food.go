package foods

type Food interface {
	Name() string
}

type Apple struct {
}

func (apple *Apple) Name() string {
	return "apple"
}
