package debug

import "github.com/pekeps/go-sc2ai/api"

var (
	white = &api.Color{R: 255, G: 255, B: 255}
	black = &api.Color{R: 0, G: 0, B: 0}

	red    = &api.Color{R: 255, G: 1, B: 1}
	blue   = &api.Color{R: 1, G: 1, B: 255}
	green  = &api.Color{R: 1, G: 255, B: 1}
	yellow = &api.Color{R: 255, G: 255, B: 1}
	orange = &api.Color{R: 255, G: 165, B: 0}
	purple = &api.Color{R: 128, G: 0, B: 128}

	lightRed    = &api.Color{R: 255, G: 102, B: 102}
	lightBlue   = &api.Color{R: 102, G: 102, B: 255}
	lightGreen  = &api.Color{R: 102, G: 255, B: 102}
	lightYellow = &api.Color{R: 255, G: 255, B: 102}
	lightOrange = &api.Color{R: 255, G: 200, B: 102}
	lightPurple = &api.Color{R: 200, G: 102, B: 200}
)
