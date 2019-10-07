package main

import "github.com/eaamj01/head-first-go/gadget"

func main() {
	player := gadget.TapePlayer{}
	mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	playList(player, mixtape)

}

func playList(device gadget.TapePlayer, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
}
