package main

import (
	"fmt"

	"github.com/eaamj01/head-first-go/gadget"
)

//Player interface
type Player interface {
	Play(string)
	Stop()
}

// TryOut Playing and recording
func TryOut(player Player) {
	player.Play("Test Track")
	player.Stop()
	// Store the tape recorder value. Use a type assertion to get a TapeRecorder value
	recorder := player.(gadget.TapeRecorder)
	// call the method that's defined on the concrete type
	recorder.Record()
}

// TryOutFixed Playing and recording
func TryOutFixed(player Player) {
	player.Play("Test Track")
	player.Stop()
	// Store the tape recorder value. Use a type assertion to get a TapeRecorder value
	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		// call the method that's defined on the concrete type
		recorder.Record()
	} else {
		fmt.Println("Player is not a TapeRecorder")
	}

}

func main() {
	TryOut(gadget.TapeRecorder{})
	// This causes a
	// panic: interface conversion: main.Player is gadget.TapePlayer, not gadget.TapeRecorder
	// TryOut(gadget.TapePlayer{})

	TryOutFixed(gadget.TapeRecorder{})
	// This causes a
	// panic: interface conversion: main.Player is gadget.TapePlayer, not gadget.TapeRecorder
	TryOutFixed(gadget.TapePlayer{})

}
