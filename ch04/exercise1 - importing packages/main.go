package main

import (
	"github.com/jaymcgavren/car"
	"github.com/jaymcgavren/car/headlights"
	"github.com/jaymcgavren/car/stereo"
	"github.com/jaymcgavren/car/wheels"
)

func main() {
	car.OpenDoor()
	headlights.TurnOn()
	stereo.TurnOn()
	stereo.BoostBass()
	wheels.Steer()
	wheels.Accelerate()

}
