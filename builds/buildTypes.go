package builds

type BuildType int

const (
	Research BuildType = iota + 1
	Train
	Build
)

func (bt BuildType) String() string {
	switch bt {
	case Research:
		return "Research"
	case Train:
		return "Train"
	case Build:
		return "Build"
	default:
		panic("Unknown build type")
	}
}
