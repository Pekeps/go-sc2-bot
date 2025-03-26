package main

import (
	"log"

	"github.com/pekeps/go-sc2ai/api"
	"github.com/pekeps/go-sc2ai/botutil"
	"github.com/pekeps/go-sc2ai/client"
	"github.com/pekeps/go-sc2ai/enums/zerg"
	"github.com/pekeps/go-sc2ai/runner"
	"github.com/pekeps/go-sc2ai/search"
)

type bot struct {
	*botutil.Bot

	myStartLocation api.Point2D
	myBases         []search.Base

	enemyStartLocation api.Point2D
	enemyBases         []search.Base

	availableBases []search.Base
}

func main() {
	// Play a random map against a medium difficulty computer
	runner.SetComputer(api.Race_Random, api.Difficulty_Easy, api.AIBuild_RandomBuild)

	// Create the agent and then start the game
	botutil.SetGameVersion()
	agent := client.AgentFunc(runAgent)
	runner.RunAgent(client.NewParticipant(api.Race_Zerg, agent, "pekeps"))
}

func runAgent(info client.AgentInfo) {
	bot := bot{Bot: botutil.NewBot(info)}
	bot.LogActionErrors()

	bot.init()

	for bot.IsInGame() {
		bot.macro()
		bot.micro()

		if err := bot.Step(1); err != nil {
			log.Print(err)
			break
		}
	}
}

func (bot *bot) init() {
	bot.initLocations()
	// Send a friendly hello
	bot.Chat("(glhf)")

}

func (bot *bot) initLocations() {
	bot.myStartLocation = bot.Self[zerg.Hatchery].First().Pos2D()
	bot.enemyStartLocation = *bot.GameInfo().StartRaw.StartLocations[0]

	log.Printf("My start location: %v", bot.myStartLocation)
	log.Printf("Enemy start location: %v", bot.enemyStartLocation)

	for _, base := range bot.availableBases {
		log.Printf("Expansion at %v", base.Location)
	}
}

func (bot *bot) macro() {

}

func (bot *bot) micro() {
}
