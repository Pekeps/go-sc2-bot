package debug

import (
	"log"

	"github.com/pekeps/go-sc2ai/api"
	"github.com/pekeps/go-sc2ai/client"
)

type GraphicalDebugger struct {
	bot     client.AgentInfo
	texts   []*api.DebugText
	lines   []*api.DebugLine
	boxes   []*api.DebugBox
	spheres []*api.DebugSphere
}

func NewGraphicalDebugger(bot client.AgentInfo) *GraphicalDebugger {
	graphicalDebugger := &GraphicalDebugger{
		bot:   bot,
		texts: []*api.DebugText{},
	}

	// Register the Draw method to be called after each step
	bot.OnAfterStep(func() {
		if bot.Observation().Observation.GameLoop%3 == 0 {
			graphicalDebugger.Draw()
		}
	})

	return graphicalDebugger
}

func (debugger *GraphicalDebugger) Draw() {
	// Draw the debug elements on the screen
	log.Printf("Drawing debug elements")
	debugger.bot.SendDebugCommands([]*api.DebugCommand{
		{
			Command: &api.DebugCommand_Draw{
				Draw: &api.DebugDraw{
					Text: debugger.texts,
				},
			},
		},
	})
}

func (d *GraphicalDebugger) AddText(text string, size uint32, virtualPos *api.Point, worldPos *api.Point, color *api.Color) {
	debugText := &api.DebugText{
		Text:       text,
		Color:      color,
		Size_:      size,
		VirtualPos: virtualPos,
		WorldPos:   worldPos,
	}
	d.texts = append(d.texts, debugText)
}

func (d *GraphicalDebugger) Clear() {
	d.texts = []*api.DebugText{}
	d.lines = []*api.DebugLine{}
	d.boxes = []*api.DebugBox{}
	d.spheres = []*api.DebugSphere{}
}
