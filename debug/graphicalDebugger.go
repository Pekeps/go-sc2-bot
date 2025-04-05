package debug

import (
	"strings"

	"github.com/pekeps/go-sc2ai/api"
	"github.com/pekeps/go-sc2ai/botutil"
	"github.com/pekeps/go-sc2ai/client"
	"github.com/pekeps/go-sc2ai/enums/ability"
	"github.com/pekeps/go-sc2ai/enums/unit"
)

// GraphicalDebugger is responsible for drawing primitive debug elements.
type GraphicalDebugger struct {
	bot     client.AgentInfo
	texts   []*api.DebugText
	lines   []*api.DebugLine
	boxes   []*api.DebugBox
	spheres []*api.DebugSphere
}

// NewGraphicalDebugger creates a new GraphicalDebugger and registers its Draw method
// to be called after each game step.
func NewGraphicalDebugger(bot client.AgentInfo) *GraphicalDebugger {
	gd := &GraphicalDebugger{
		bot: bot,
	}
	// Register the Draw callback so that debug commands are sent every step.
	bot.OnAfterStep(func() {
		if len(gd.texts) == 0 && len(gd.lines) == 0 && len(gd.boxes) == 0 && len(gd.spheres) == 0 {
			println("Nothing to draw")
		} else {
			println("drawing")
		}
		gd.Draw()

	})
	return gd
}

// Draw sends the accumulated debug commands (text, lines, boxes, spheres) to the bot.
func (gd *GraphicalDebugger) Draw() {
	gd.bot.SendDebugCommands([]*api.DebugCommand{
		{
			Command: &api.DebugCommand_Draw{
				Draw: &api.DebugDraw{
					Text:    gd.texts,
					Lines:   gd.lines,
					Boxes:   gd.boxes,
					Spheres: gd.spheres,
				},
			},
		},
	})
	gd.Clear()
}

// Clear resets all debug elements.
func (gd *GraphicalDebugger) Clear() {
	gd.texts = []*api.DebugText{}
	gd.lines = []*api.DebugLine{}
	gd.boxes = []*api.DebugBox{}
	gd.spheres = []*api.DebugSphere{}
}

// AddText adds a new debug text element.
func (gd *GraphicalDebugger) AddText(text string, size uint32, virtualPos *api.Point, worldPos *api.Point, color *api.Color) {
	dt := &api.DebugText{
		Text:       text,
		Color:      color,
		Size_:      size,
		VirtualPos: virtualPos,
		WorldPos:   worldPos,
	}
	gd.texts = append(gd.texts, dt)
}

// AddSphere adds a debug sphere at the given position.
func (gd *GraphicalDebugger) AddSphere(pos *api.Point, radius float32, color *api.Color) {
	ds := &api.DebugSphere{
		Color: color,
		P:     pos,
		R:     radius,
	}
	gd.spheres = append(gd.spheres, ds)
}

// AddLine adds a debug line from start to end.
func (gd *GraphicalDebugger) AddLine(start *api.Point, end *api.Point, color *api.Color) {
	dl := &api.DebugLine{
		Color: color,
		Line: &api.Line{
			P0: start,
			P1: end,
		},
	}
	gd.lines = append(gd.lines, dl)
}

// AddUnit draws a unit's sphere and label, and optionally its orders.
func (gd *GraphicalDebugger) AddUnit(u *botutil.Unit, color *api.Color, drawOrders bool) {
	if color == nil {
		color = unitColor(u)
	}
	// Use a part of the unit type string as its label.
	parts := strings.Split(unit.String(u.GetUnitType()), "_")
	unitName := ""
	if len(parts) > 1 {
		unitName = parts[1]
	} else {
		unitName = unit.String(u.GetUnitType())
	}
	gd.AddSphere(u.Pos, u.Radius, color)
	gd.AddText(unitName, 8, nil, u.Pos, color)
	if drawOrders {
		gd.drawOrders(u)
	}
}

// drawOrders draws lines representing the orders of the given unit.
func (gd *GraphicalDebugger) drawOrders(u *botutil.Unit) {
	orders := u.GetOrders()
	for _, order := range orders {
		abilityId := ability.Remap(order.GetAbilityId())
		target := order.GetTargetWorldSpacePos()
		var targetPoint *api.Point
		if target != nil {
			target.Z = u.Pos.Z // keep the target at the same height as the unit
			targetDistance := u.Pos.Distance(*target)
			offsetPoint := u.Pos.Offset(*target, targetDistance)
			targetPoint = &offsetPoint
		}
		switch abilityId {
		case ability.Move:
			gd.AddLine(u.Pos, targetPoint, white)
		case ability.Attack:
			gd.AddLine(u.Pos, targetPoint, red)
		case ability.Harvest_Gather, ability.Harvest_Return:
			gd.AddLine(u.Pos, targetPoint, blue)
		}
	}
}

// unitColor returns a default color for a unit based on its alliance.
func unitColor(u *botutil.Unit) *api.Color {
	switch u.Alliance {
	case api.Alliance_Self:
		return lightGreen
	case api.Alliance_Enemy:
		return lightRed
	case api.Alliance_Neutral:
		return lightYellow
	default:
		return lightPurple
	}
}
