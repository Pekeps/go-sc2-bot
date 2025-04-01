package managers

import (
	"log"

	"github.com/pekeps/go-sc2ai/api"
	"github.com/pekeps/go-sc2ai/botutil"
	"github.com/pekeps/go-sc2ai/enums/ability"
	"github.com/pekeps/go-sc2ai/enums/zerg"
)

type EconomyManager struct {
	bot *botutil.Bot

	minerals  uint32
	gas       uint32
	larvae    botutil.Units
	supply    uint32
	supplyCap uint32
	workers   botutil.Units
}

func NewEconomyManager(bot *botutil.Bot) *EconomyManager {
	manager := &EconomyManager{bot: bot}
	manager.Init()
	return manager
}

func (m *EconomyManager) Init() {

}

func (m *EconomyManager) Manage(loop uint32) {
	m.update(m.bot.Observation().Observation.PlayerCommon)

	if m.getSupplyLeft() < 2 && m.larvae.Len() > 0 {
		if m.bot.BuildUnit(zerg.Larva, ability.Train_Overlord) {
			return
		}
	}

	m.bot.BuildUnit(zerg.Larva, ability.Train_Drone)
}

func (m *EconomyManager) Destruct() {

}

func (m *EconomyManager) update(pc *api.PlayerCommon) {
	m.minerals = pc.GetMinerals()
	m.gas = pc.GetVespene()
	m.larvae = m.bot.Self[zerg.Larva]
	m.supply = pc.GetFoodUsed()
	m.supplyCap = pc.GetFoodCap()

	m.workers = m.bot.Self[zerg.Drone]
}

func (m *EconomyManager) getSupplyLeft() uint32 {
	return m.supplyCap - m.supply
}

func (m *EconomyManager) printDebug() {
	log.Printf("EconomyManager")
	log.Printf("Minerals: %d", m.minerals)
	log.Printf("Gas: %d", m.gas)
	log.Printf("Supply: %d/%d", m.supply, m.supplyCap)
	log.Printf("Larvae: %d", m.larvae.Len())
	log.Printf("Workers: %d", m.workers.Len())
}
