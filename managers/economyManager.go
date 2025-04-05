package managers

import (
	"github.com/pekeps/go-sc2ai/botutil"
	"github.com/pekeps/go-sc2ai/builds"
)

type EconomyManager struct {
	player *botutil.Player
}

func NewEconomyManager(player *botutil.Player) *EconomyManager {
	manager := &EconomyManager{player}
	manager.Init()
	return manager
}

func (m *EconomyManager) Init() {

}

func (m *EconomyManager) Manage() {

}

func (m *EconomyManager) GetBuildRequests() []*builds.BuildRequest {
	return []*builds.BuildRequest{}
}
