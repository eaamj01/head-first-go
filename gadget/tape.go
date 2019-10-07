package gadget

import "fmt"

// TapePlayer has batteries
type TapePlayer struct {
	Batteries string
}

// Play Song on Tape Player
func (t TapePlayer) Play(song string) {
	fmt.Println("Playing", song)
}

// Stop Tape Player
func (t TapePlayer) Stop() {
	fmt.Println("Stopped")
}

// TapeRecorder has a mic
type TapeRecorder struct {
	Microphone int
}

// Play Song on Tape Recorder
func (t TapeRecorder) Play(song string) {
	fmt.Println("Playing", song)
}

// Record Tape Recorder
func (t TapeRecorder) Record() {
	fmt.Println("Recording")
}

// Stop Tape Recorder
func (t TapeRecorder) Stop() {
	fmt.Println("Stopped")
}
