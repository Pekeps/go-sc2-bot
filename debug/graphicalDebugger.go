package debug

import (
	"strings"

	"github.com/pekeps/go-sc2ai/api"
	"github.com/pekeps/go-sc2ai/client"
	"github.com/pekeps/go-sc2ai/enums/ability"
	"github.com/pekeps/go-sc2ai/enums/unit"
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
		bot: bot,
	}

	// Register the Draw method to be called after each step
	bot.OnAfterStep(func() {
		graphicalDebugger.Draw()
	})

	return graphicalDebugger
}

func (debugger *GraphicalDebugger) Draw() {
	// Draw the debug elements on the screen
	debugger.bot.SendDebugCommands([]*api.DebugCommand{
		{
			Command: &api.DebugCommand_Draw{
				Draw: &api.DebugDraw{
					Text:    debugger.texts,
					Lines:   debugger.lines,
					Boxes:   debugger.boxes,
					Spheres: debugger.spheres,
				},
			},
		},
	})
}

func (d *GraphicalDebugger) Clear() {
	d.texts = []*api.DebugText{}
	d.lines = []*api.DebugLine{}
	d.boxes = []*api.DebugBox{}
	d.spheres = []*api.DebugSphere{}
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

func (d *GraphicalDebugger) AddSphere(pos *api.Point, radius float32, color *api.Color) {
	debugSphere := &api.DebugSphere{
		Color: color,
		P:     pos,
		R:     radius,
	}
	d.spheres = append(d.spheres, debugSphere)
}

func (d *GraphicalDebugger) AddLine(start *api.Point, end *api.Point, color *api.Color) {
	debugLine := &api.DebugLine{
		Color: color,
		Line: &api.Line{
			P0: start,
			P1: end,
		},
	}
	d.lines = append(d.lines, debugLine)
}

func (g *GraphicalDebugger) AddUnit(u *api.Unit, color *api.Color, drawOrders bool) {
	if color == nil {
		color = unitColor(u)
	}
	// First part is always race prefix
	unitName := strings.Split(unit.String(u.GetUnitType()), "_")[1]

	g.AddSphere(u.Pos, u.Radius, color)
	g.AddText(unitName, 8, nil, u.Pos, color)
	if drawOrders {
		g.drawOrders(u)
	}
}
func (g *GraphicalDebugger) drawOrders(u *api.Unit) {
	orders := u.GetOrders()
	for _, order := range orders {
		abilityId := ability.Remap(order.GetAbilityId())
		target := order.GetTargetWorldSpacePos()
		var targetPoint *api.Point = nil
		if target != nil {
			target.Z = u.Pos.Z
			targetDistance := u.Pos.Distance(*target)
			offsetPoint := u.Pos.Offset(*target, targetDistance)
			targetPoint = &offsetPoint
		}
		if abilityId == ability.Move {
			g.AddLine(u.Pos, targetPoint, white)
		}
		if abilityId == ability.Attack {
			g.AddLine(u.Pos, targetPoint, red)
		}

		// TODO Base and resource target is nil
		if abilityId == ability.Harvest_Gather || abilityId == ability.Harvest_Return {
			g.AddLine(u.Pos, targetPoint, blue)
		}
	}
}

func unitColor(unit *api.Unit) *api.Color {
	alliance := unit.Alliance
	if alliance == api.Alliance_Self {
		return lightGreen
	}
	if alliance == api.Alliance_Enemy {
		return lightRed
	}
	if alliance == api.Alliance_Neutral {
		return lightYellow
	}

	return lightPurple
}
