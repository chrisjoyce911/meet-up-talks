package strawman

import (
	"fmt"
	"math/rand"
	"time"
)

var l float32

func (s *Strawman) MakeHay() {

	fmt.Println("Making Hay")
	time.Sleep(2 * time.Second)

	l = getvalue()

	for i := 1; i < 25; i++ {

		n := getvalue()

		if n < l {
			if s.quantity > 0 {
				s.quantity = s.quantity - 1
				fmt.Printf("Selling at\t: %s$%f\n", s.currency, float32(1+getvalue()))
			}
		} else {
			s.quantity = s.quantity + 4
			fmt.Printf("Buying at\t: %s$%f\n", s.currency, float32(1+getvalue()))
		}
		l = n
		time.Sleep(1 * time.Second)
	}

}

func (s *Strawman) BuyUnits(n int) {
	s.quantity = s.quantity + n
	fmt.Printf("Balance : %d\n", s.quantity)
}

func (s *Strawman) CurrentUnits() {
	fmt.Printf("Curret Balance\t: %d\n", s.quantity)
}

func (s *Strawman) CurrentValue() {
	fmt.Printf("Current Value\t: %s$%f\n", s.currency, (float32(s.quantity)*1 + getvalue()))
}

func getvalue() float32 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float32()
}
