package managers

import (
	"slices"

	"github.com/pekeps/go-sc2ai/botutil"
	"github.com/pekeps/go-sc2ai/builds"
	"github.com/pekeps/go-sc2ai/enums/ability"
)

type BuildManager struct {
	player             *botutil.Player
	buildOrder         builds.BuildOrder
	buildOrderComplete bool

	buildRequests []*builds.BuildRequest
}

func NewBuildManager(buildOrder builds.BuildOrder, player *botutil.Player) *BuildManager {
	return &BuildManager{
		player:             player,
		buildOrder:         buildOrder,
		buildOrderComplete: false,
	}
}

func (bm *BuildManager) Init() {
	bm.buildRequests = bm.buildOrder.GetAllSteps()
}

func (bm *BuildManager) Manage() {
	bm.pruneBuildRequests()

	if !bm.buildOrderComplete {
		uncompletedSteps := bm.buildOrder.GetUncompletedSteps()
		availableBuildSteps := slices.ContainsFunc(uncompletedSteps, func(step *builds.BuildRequest) bool {
			return bm.player.GetFoodUsed() >= uint32(step.GetAtSupply())
		})
		if !availableBuildSteps && bm.player.CanAfford(botutil.Cost{Minerals: 50}) {
			bm.buildRequests = append(bm.buildRequests, builds.NewTrainOrBuildRequest(builds.Train, ability.Train_Drone, 1, builds.Normal, 0))
		}

	}
	if len(bm.buildOrder.GetUncompletedSteps()) == 0 {
		bm.buildOrderComplete = true
	}

}

func (bm *BuildManager) GetBuildRequests() []*builds.BuildRequest {
	return bm.buildRequests
}

func (bm *BuildManager) pruneBuildRequests() {
	for i := 0; i < len(bm.buildRequests); {
		if bm.buildRequests[i].IsCompleted() {
			bm.buildRequests = slices.Delete(bm.buildRequests, i, i+1)
		} else {
			i++
		}
	}
}
