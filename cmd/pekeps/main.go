package main

import (
	"log"

	"github.com/pekeps/go-sc2ai/api"
	"github.com/pekeps/go-sc2ai/botutil"
	"github.com/pekeps/go-sc2ai/client"
	"github.com/pekeps/go-sc2ai/debug"
	"github.com/pekeps/go-sc2ai/managers"
	"github.com/pekeps/go-sc2ai/runner"
	"github.com/pekeps/go-sc2ai/search"
)

type bot struct {
	*botutil.Bot
	*search.Map
	*debug.Debugger

	*managers.Hub

	loop uint32
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
		bot.loop++
		bot.macro()
		bot.micro()
		if err := bot.Step(1); err != nil {
			log.Print(err)
			break
		}
	}

	bot.destruct()
}

func (bot *bot) init() {
	log.Printf("Initializing pekeps bot")
	bot.Hub = managers.NewHub(bot.Bot)
	bot.Debugger = debug.NewDebugger(bot.AgentInfo)

	bot.Chat("(glhf)")

}

func (bot *bot) destruct() {
	log.Printf("Pekeps bot destructing")
}
