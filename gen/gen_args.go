package gen

type GenConfiguration struct {
	Depth uint
}

type GenOption func(args *GenConfiguration)

func WithDepth(d uint) GenOption {
	return func(gc *GenConfiguration) {
		gc.Depth = d
	}
}
