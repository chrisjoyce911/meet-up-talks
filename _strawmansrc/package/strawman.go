package strawman

type Strawman struct {
	currency  string // HL
	transport string // HL
	quantity  int    // HL
	use       string // HL
}

// New provides access to a new straw trading calculator
func New(currency string, opts ...Option) *Strawman {
	straw := &Strawman{
		currency:  currency, // HL
		quantity:  1,        // HL
		use:       "feed",   // HL
		transport: "bale",   // HL
	}

	for _, applyOpt := range opts {
		applyOpt(straw) // HL
	}

	return straw
}
