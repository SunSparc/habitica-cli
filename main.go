package main

import (
	"habitica-cli/client"
)

func main() {
	// TODO: accept flags (eg rotateCurrentMount, hatchEggs, feedPets, maxDevUserAccount)
	// Flags
	//   - mode: dev, live (default=dev)
	whatToDo := "determined by flags or defaults"

	// todo: receive mode via flag/env/file, default to "dev"
	habiticaClient := client.NewHabiticaApiClient("live")

	live := client.NewLiveMode(habiticaClient.HabiticaUser)
	live.ManageMounts()

	// TODO: how do we choose what to do? switch flags?
	switch whatToDo {
	case "":
		// do it
	default:
		// do that
	}
}
