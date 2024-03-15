package app

type Interface interface {
}

type app struct {
}

func NewApp() Interface {
	a := app{}
	return &a
}
