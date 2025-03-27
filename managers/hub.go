package managers

import "github.com/pekeps/go-sc2ai/botutil"

type Hub struct {
	*botutil.Bot

	Managers *[]manager
}

func NewHub(bot *botutil.Bot) *Hub {
	managers := []manager{
		NewEconomyManager(bot),
	}
	return &Hub{Bot: bot, Managers: &managers}
}

func (h *Hub) StepManagers(loop uint32) {
	for _, m := range *h.Managers {
		m.Manage(loop)
	}
}

// func initManagers(bot *botutil.Bot) *[]manager {
// 	managers := []manager{
// 		NewEconomyManager(bot),
// 	}
// 	return &managers
// }
