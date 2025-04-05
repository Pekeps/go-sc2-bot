package debug

import "github.com/pekeps/go-sc2ai/api"

// DrawOptions holds options for drawing text.
type DrawOptions struct {
	Size       uint32
	Color      *api.Color
	virtualPos *api.Point
	worldPos   *api.Point
}

// Option is a function that configures DrawOptions.
type Option func(*DrawOptions)

// WithVirtualPos sets a custom virtual position.
func WithVirtualPos(pos *api.Point) Option {
	return func(do *DrawOptions) {
		do.virtualPos = pos
	}
}

// WithWorldPos sets a custom world position.
func WithWorldPos(pos *api.Point) Option {
	return func(do *DrawOptions) {
		do.worldPos = pos
	}
}

// WithSize sets a custom text size.
func WithSize(size uint32) Option {
	return func(do *DrawOptions) {
		do.Size = size
	}
}

// WithColor sets a custom text color.
func WithColor(color *api.Color) Option {
	return func(do *DrawOptions) {
		do.Color = color
	}
}
