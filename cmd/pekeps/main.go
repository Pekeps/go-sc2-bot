package main

import (
	"fmt"
	"log"

	"github.com/pekeps/go-sc2ai/api"
	"github.com/pekeps/go-sc2ai/botutil"
	"github.com/pekeps/go-sc2ai/builds"
	"github.com/pekeps/go-sc2ai/client"
	"github.com/pekeps/go-sc2ai/debug"
	"github.com/pekeps/go-sc2ai/managers"
	"github.com/pekeps/go-sc2ai/runner"
	"github.com/pekeps/go-sc2ai/search"
)

type bot struct {
	*botutil.Bot

	gameMap  *search.Map
	debugger *debug.Debugger
	managers []managers.Manager
}

func main() {
	// Play a random map against a medium difficulty computer
	runner.SetComputer(api.Race_Random, api.Difficulty_Easy, api.AIBuild_RandomBuild)

	// Create the agent and then start the game
	botutil.SetGameVersion()
	var agent client.AgentFunc = runAgent
	runner.RunAgent(client.NewParticipant(api.Race_Zerg, agent, "pekeps"))
}

func runAgent(info client.AgentInfo) {
	bot := bot{Bot: botutil.NewBot(info)}
	bot.LogActionErrors()

	bot.init()

	for bot.IsInGame() {
		bot.debugger.DebugPanel.AddEntry(fmt.Sprintf("Supply %d/%d", bot.Player.GetFoodUsed(), bot.Player.GetFoodCap()))
		bot.debugger.DebugPanel.AddEntry("")
		bot.executeManagers()
		buildRequests := bot.getBuildRequests()
		bot.handleBuildRequests(buildRequests)

		bot.debugger.Draw()
		if err := bot.Step(1); err != nil {
			log.Print(err)
			break
		}
	}

	bot.destruct()
}

func (bot *bot) init() {
	log.Printf("Initializing pekeps bot")
	bot.gameMap = search.NewMap(bot.Bot)
	bot.debugger = debug.NewDebugger(bot.Bot, bot.UnitContext)

	bot.managers = []managers.Manager{
		managers.NewEconomyManager(bot.Player),
		managers.NewBuildManager(builds.LingFlood(), bot.Player),
	}

	for _, manager := range bot.managers {
		manager.Init()
	}

	bot.Chat("(glhf)")

}

func (bot *bot) destruct() {
	log.Printf("Pekeps bot destructing")
}

func (bot *bot) executeManagers() {
	for _, manager := range bot.managers {
		manager.Manage()
	}
}

func (bot *bot) getBuildRequests() []*builds.BuildRequest {
	var requests []*builds.BuildRequest
	for _, manager := range bot.managers {
		requests = append(requests, manager.GetBuildRequests()...)
	}
	return requests
}
