package botutil

import (
	"github.com/pekeps/go-sc2ai/api"
)

type Base struct {
	Base     *api.Unit
	Location api.Point2D
	Minerals []Unit
	Vespene  []Unit

	// Distance from starter base
	Distance float32
}
