package debug

import (
	"slices"

	"github.com/pekeps/go-sc2ai/api"
)

// DebugPanel manages debug text entries that are drawn as an overlay.
type DebugPanel struct {
	entries []string
}

// AddEntry adds a new text entry to the panel.
func (dp *DebugPanel) AddEntry(text string) {
	dp.entries = append(dp.entries, text)
}

// RemoveEntry removes the entry at the given index.
func (dp *DebugPanel) RemoveEntry(index int) {
	if index >= 0 && index < len(dp.entries) {
		dp.entries = slices.Delete(dp.entries, index, index+1)
	}
}

// Clear removes all entries from the panel.
func (dp *DebugPanel) Clear() {
	dp.entries = []string{}
}

// GetEntries returns all entries in the panel.
func (dp *DebugPanel) GetEntries() []string {
	return dp.entries
}

// Draw renders each debug panel entry using the provided callback.
// The callback should draw text at the provided virtual position.
func (dp *DebugPanel) Draw(drawTextFunc func(text string, virtualPos *api.Point)) {
	const (
		baseX      = 0.05
		baseY      = 0.05
		lineHeight = 0.015
	)
	for i, entry := range dp.entries {
		virtualPos := &api.Point{
			X: baseX,
			Y: baseY + float32(i)*lineHeight,
			Z: 0,
		}
		drawTextFunc(entry, virtualPos)
	}
}
