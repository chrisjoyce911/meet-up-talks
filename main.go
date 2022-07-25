package main

import "talk-optional-parameters/strawman"

func main() {
	straw := strawman.New("important")
	straw = strawman.New("important", strawman.WithNum(30))
	straw = strawman.New("required", strawman.WithNum(20), strawman.WithStr("hello"))

	straw.MakeHay()

}
