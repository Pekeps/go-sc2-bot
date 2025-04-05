package builds

import (
	"github.com/pekeps/go-sc2ai/enums/ability"
	"github.com/pekeps/go-sc2ai/enums/upgrade"
)

// "builds/buildOrdersnewBuildOrder" // Removed or corrected if necessary
type BuildOrder struct {
	name  string
	steps []*BuildRequest
}

func (bo BuildOrder) String() string {
	return bo.name
}

func (bo BuildOrder) GetAllSteps() []*BuildRequest {
	return bo.steps
}

func (bo BuildOrder) GetUncompletedSteps() []*BuildRequest {
	var uncompletedSteps []*BuildRequest
	for _, step := range bo.steps {
		if !step.IsCompleted() {
			uncompletedSteps = append(uncompletedSteps, step)
		}
	}
	return uncompletedSteps
}

// https://lotv.spawningtool.com/build/188392/

func LingFlood() BuildOrder {
	return BuildOrder{
		name: "Ling Flood",
		steps: []*BuildRequest{
			NewTrainOrBuildRequest(Train, ability.Train_Overlord, 1, Normal, 13),
			NewTrainOrBuildRequest(Build, ability.Build_Hatchery, 1, Normal, 17),
			NewTrainOrBuildRequest(Build, ability.Build_Extractor, 1, Normal, 18),
			NewTrainOrBuildRequest(Build, ability.Build_SpawningPool, 1, Normal, 17),
			NewTrainOrBuildRequest(Train, ability.Train_Overlord, 1, Normal, 19),
			NewTrainOrBuildRequest(Train, ability.Train_Queen, 2, Normal, 20),
			NewTrainOrBuildRequest(Train, ability.Train_Zergling, 2, Normal, 24),
			NewResearchRequest(ability.Research_ZerglingMetabolicBoost, 1, Normal, 28, upgrade.Zerglingmovementspeed),
			NewTrainOrBuildRequest(Build, ability.Build_Hatchery, 1, Normal, 32),
			NewTrainOrBuildRequest(Train, ability.Train_Overlord, 1, Normal, 31),
		},
	}

}
