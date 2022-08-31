package strawman

// OPTIONSSTART OMIT
type Option func(s *Strawman) // HL

// Quantity of transport units
func Quantity(num int) Option { // HL
	return func(s *Strawman) {
		s.quantity = num
	}
}

// UsedFor is the intended use of the staw
func UsedFor(str string) Option { // HL
	return func(s *Strawman) {
		s.use = str
	}
}

// OPTIONSEND OMIT

// CUSTOMSTART OMIT
// ForFeed as part of the roughage component of the diet to cattle or horses
func ForFeed() Option { // HL
	return func(s *Strawman) {
		s.use = "feed"
		s.transport = "bale"
	}
}

// ForMulch to cover the garden beds
func ForMulch() Option { // HL
	return func(s *Strawman) {
		s.use = "mulch"
		s.transport = "bale"
	}
}

// ForBedding to be used in straw-filled mattress, also known as a palliasse.
func ForBedding() Option { // HL
	return func(s *Strawman) {
		s.use = "bedding"
		s.transport = "bag"
	}
}

// CUSTOMEND OMIT
