package main

import strawman "github.com/chrisjoyce911/meet-up-talks/strawman/package"

func main() {
	straw := strawman.New("AUD")
	straw = strawman.New("AUD", strawman.Quantity(30))
	straw = strawman.New("AUD", strawman.Quantity(30), strawman.UsedFor("barskets"))

	// Using a preset
	straw = strawman.New("AUD", strawman.ForFeed())

	straw.BuyUnits(10)

	straw.CurrentUnits()
	straw.CurrentValue()

	straw.MakeHay()

	straw.CurrentUnits()
	straw.CurrentValue()

}
