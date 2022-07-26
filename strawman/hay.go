package strawman

import (
	"log"
	"math/rand"
	"time"
)

var l float32

func (s *Strawman) MakeHay() {

	log.Print("Making Hay")
	time.Sleep(2 * time.Second)

	l = getvalue()

	for i := 1; i < 25; i++ {

		n := getvalue()

		if n < l {
			if s.quantity > 0 {
				s.quantity = s.quantity - 1
				log.Printf("Selling at\t: %s$%f", s.currency, float32(1+getvalue()))
			}
		} else {
			s.quantity = s.quantity + 1
			log.Printf("Buying at\t: %s$%f", s.currency, float32(1+getvalue()))
		}
		time.Sleep(1 * time.Second)
	}

}

func (s *Strawman) BuyUnits(n int) {
	s.quantity = s.quantity + n
	log.Printf("Balance : %d", s.quantity)
}

func (s *Strawman) CurrentUnits() {
	log.Printf("Curret Balance\t: %d", s.quantity)
}

func (s *Strawman) CurrentValue() {
	log.Printf("Current Value\t: %s$%f", s.currency, (float32(s.quantity)*1 + getvalue()))
}

func getvalue() float32 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float32()
}
