package main

import (
	"github.com/pekeps/go-sc2ai/botutil"
	"github.com/pekeps/go-sc2ai/builds"
	"github.com/pekeps/go-sc2ai/enums/zerg"
)

func (bot *bot) handleBuildRequests(requests []*builds.BuildRequest) {
	for _, request := range requests {
		bot.debugger.DebugPanel.AddEntry(request.String())

		if bot.Player.GetFoodUsed() < uint32(request.GetAtSupply()) {
			continue
		}

		switch request.GetBuildType() {
		case builds.Build:
			build(request, bot.Builder)
		case builds.Train:
			train(request, bot.Builder)
		case builds.Research:
			research(request, bot.Builder)
		default:
			panic("Unknown build action")
		}
	}
}

// build returns
func build(request *builds.BuildRequest, builder *botutil.Builder) int {
	fulfilled := builder.BuildUnits(zerg.Drone, request.GetAbilityID(), int(request.GetUnfulfilled()))
	request.Fulfill(fulfilled)
	return fulfilled
}

func train(request *builds.BuildRequest, builder *botutil.Builder) int {
	fulfilled := builder.BuildUnits(zerg.Larva, request.GetAbilityID(), int(request.GetUnfulfilled()))
	request.Fulfill(fulfilled)
	return fulfilled
}

func research(request *builds.BuildRequest, builder *botutil.Builder) bool {
	fulfilled := builder.BuildUnit(zerg.EvolutionChamber, request.GetAbilityID())
	if fulfilled {
		request.Fulfill(1)
	}
	return fulfilled
}
