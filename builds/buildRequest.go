package builds

import (
	"strings"

	"fmt"

	"github.com/pekeps/go-sc2ai/api"
	"github.com/pekeps/go-sc2ai/cmd/gen_ids/enums/ability"
)

type BuildRequest struct {
	buildType BuildType
	abilityID api.AbilityID
	upgradeID api.UpgradeID
	amount    int
	fulfilled int
	priority  BuildPriority
	atSupply  int
}

func NewTrainOrBuildRequest(buildType BuildType, command api.AbilityID, amount int, priority BuildPriority, atSupply int) *BuildRequest {
	return newBuildRequest(buildType, command, amount, priority, atSupply, 0)
}

func NewResearchRequest(command api.AbilityID, amount int, priority BuildPriority, atSupply int, researchID api.UpgradeID) *BuildRequest {
	return newBuildRequest(Research, command, amount, priority, atSupply, researchID)
}

func newBuildRequest(buildType BuildType, command api.AbilityID, amount int, priority BuildPriority, atSupply int, researchID api.UpgradeID) *BuildRequest {
	return &BuildRequest{
		buildType: buildType,
		abilityID: command,
		upgradeID: researchID,
		amount:    amount,
		fulfilled: 0,
		priority:  priority,
		atSupply:  atSupply,
	}
}

func (br *BuildRequest) String() string {
	action := strings.ReplaceAll(ability.String(br.abilityID), "_", " ")
	amount := fmt.Sprint(br.fulfilled) + "/" + fmt.Sprint(br.amount)

	if br.atSupply > 0 {
		amount += " at " + fmt.Sprint(br.atSupply)
	}

	return action + " " + amount
}

func (br *BuildRequest) Fulfill(i int) {
	br.fulfilled += i
}

func (br *BuildRequest) GetUnfulfilled() int {
	return br.amount - br.fulfilled
}

func (br *BuildRequest) IsCompleted() bool {
	return br.GetUnfulfilled() <= 0
}

func (br *BuildRequest) GetBuildType() BuildType {
	return br.buildType
}

func (br *BuildRequest) GetAbilityID() api.AbilityID {
	return br.abilityID
}

func (br *BuildRequest) GetUpgradeID() api.UpgradeID {
	return br.upgradeID
}

func (br *BuildRequest) GetAmount() int {
	return br.amount
}

func (br *BuildRequest) GetFulfilled() int {
	return br.fulfilled
}

func (br *BuildRequest) GetPriority() BuildPriority {
	return br.priority
}

func (br *BuildRequest) GetAtSupply() int {
	return br.atSupply
}
