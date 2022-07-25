package strawman

type Option func(s *Strawman)

func WithNum(num int) Option {
	return func(s *Strawman) {
		s.Num = num
	}
}

func WithStr(str string) Option {
	return func(s *Strawman) {
		s.Str = str
	}
}
