package main

import (
	"example.com/pronist/gadget"
)

func playList(device gadget.Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func TryOut(player gadget.Player) {
	player.Play("Test Track")
	player.Stop()

	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
	}
}

func main() {
	var player gadget.Player

	player = gadget.TapePlayer{}
	mixtape := []string{"Jessie's Girl", "Whip It"}
	playList(player, mixtape)

	player = gadget.TapeRecorder{}
	playList(player, mixtape)

	TryOut(gadget.TapePlayer{})
	TryOut(gadget.TapeRecorder{})
}
