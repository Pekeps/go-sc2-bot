package main

import (
	"log"

	"github.com/pekeps/go-sc2ai/api"
	"github.com/pekeps/go-sc2ai/botutil"
	"github.com/pekeps/go-sc2ai/client"
	"github.com/pekeps/go-sc2ai/runner"
)

func main() {
	// Play a random map against a medium difficulty computer
	runner.SetComputer(api.Race_Random, api.Difficulty_Medium, api.AIBuild_RandomBuild)

	// Create the agent and then start the game
	botutil.SetGameVersion()
	log.Printf("Set game version to %v", botutil.GameVersion)
	agent := client.AgentFunc(runAgent)
	log.Printf("Starting agent")
	runner.RunAgent(client.NewParticipant(api.Race_Zerg, agent, "ZergRush"))
}
