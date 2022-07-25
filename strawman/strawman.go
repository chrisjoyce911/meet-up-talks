package strawman

type Strawman struct {
	Num int
	Str string
}

func New(someRequiredField string, opts ...Option) *Strawman {
	straw := &Strawman{
		Num: 10,
		Str: "hello",
	}

	for _, applyOpt := range opts {
		applyOpt(straw)
	}

	return straw
}
