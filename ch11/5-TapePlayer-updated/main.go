package main

import "github.com/eaamj01/head-first-go/gadget"

// Player interface
type Player interface {
	Play(string)
	Stop()
}

func main() {
	mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	var player Player = gadget.TapePlayer{}
	playList(player, mixtape)

}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}
