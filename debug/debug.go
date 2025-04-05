package debug

import (
	"strings"

	"github.com/pekeps/go-sc2ai/api"
	"github.com/pekeps/go-sc2ai/botutil"
	"github.com/pekeps/go-sc2ai/client"
)

// Debugger integrates both the graphical debugger and debug panel.
type Debugger struct {
	graphics    *GraphicalDebugger
	DebugPanel  *DebugPanel
	unitContext *botutil.UnitContext
}

// NewDebugger creates a new Debugger instance.
func NewDebugger(bot client.AgentInfo, unitContext *botutil.UnitContext) *Debugger {
	return &Debugger{
		graphics:    NewGraphicalDebugger(bot),
		DebugPanel:  &DebugPanel{},
		unitContext: unitContext,
	}
}

// DrawText renders a single line of debug text using the provided options.
func (d *Debugger) DrawText(text string, opts ...Option) {
	// Set default drawing options.
	options := &DrawOptions{
		Size:       12,
		Color:      lightYellow,
		virtualPos: &api.Point{X: 0.05, Y: 0.05, Z: 0},
		worldPos:   nil,
	}
	// Apply any custom options.
	for _, opt := range opts {
		opt(options)
	}
	d.graphics.AddText(text, options.Size, options.virtualPos, options.worldPos, options.Color)
}

// DrawTextSlice joins multiple lines of text and draws them.
func (d *Debugger) DrawTextSlice(texts []string, opts ...Option) {
	d.DrawText(strings.Join(texts, "\n"), opts...)
}

// DrawUnits renders debug information for all units provided in the unit context.
func (d *Debugger) DrawUnits(ctx *botutil.UnitContext) {
	// Iterate over all units and render each one if visible.
	ctx.AllUnits().Each(func(u botutil.Unit) {
		if u.IsVisible() {
			d.graphics.AddUnit(&u, nil, true)
		}
	})
}

// Draw is the primary method called each frame to render all debug elements.
func (d *Debugger) Draw() {
	// Draw all units
	d.DrawUnits(d.unitContext)

	// Draw the debug panel entries.
	d.DebugPanel.Draw(func(text string, pos *api.Point) {
		d.DrawText(text, WithVirtualPos(pos))
	})
	// Clear debug panel entries after drawing so they are not drawn repeatedly.
	d.DebugPanel.Clear()
}
