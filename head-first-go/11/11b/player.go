// B"H

package main

import "github.com/Ylazerson/go-shenanigans/head-first-go/11/11b/gadget"

// -- -------------------------------------
type Player interface {
	Play(string)
	Stop()
}

// -- -------------------------------------
func TryOut(player Player) {

	player.Play("Test Track")

	player.Stop()

	// -- ---------------------------------
	// type assertion:
	recorder, ok := player.(gadget.TapeRecorder)

	if ok {
		recorder.Record()
	}
}

// -- -------------------------------------
func main() {
	TryOut(gadget.TapeRecorder{})
	TryOut(gadget.TapePlayer{})
}
