package debug

import (
	"strings"

	"github.com/pekeps/go-sc2ai/api"
	"github.com/pekeps/go-sc2ai/client"
)

type Debugger struct {
	graphics *GraphicalDebugger
}

func NewDebugger(bot client.AgentInfo) *Debugger {
	return &Debugger{
		graphics: NewGraphicalDebugger(bot),
	}
}

type DrawOptions struct {
	Size       uint32
	Color      *api.Color
	virtualPos *api.Point
	worldPos   *api.Point
}
type Option func(*DrawOptions)

func (debugger *Debugger) DrawText(text string, opts ...Option) {
	options := &DrawOptions{
		Size:  12,
		Color: lightYellow,
		virtualPos: &api.Point{
			X: 0.05,
			Y: 0.05,
			Z: 0,
		},
		worldPos: nil,
	}

	for _, opt := range opts {
		opt(options)
	}

	debugger.graphics.AddText(text, options.Size, options.virtualPos, options.worldPos, options.Color)
}

func (debugger *Debugger) DrawTextSlice(texts []string, opts ...Option) {
	debugger.DrawText(strings.Join(texts, "\n"), opts...)
}
