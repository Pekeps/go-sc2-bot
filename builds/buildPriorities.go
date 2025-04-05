package builds

type BuildPriority int

const (
	VeryLow BuildPriority = iota + 1
	Low
	Normal
	High
	VeryHigh
)

func (bp BuildPriority) String() string {
	switch bp {
	case VeryLow:
		return "Very Low"
	case Low:
		return "Low"
	case Normal:
		return "Normal"
	case High:
		return "High"
	case VeryHigh:
		return "Very High"
	default:
		return "Unknown"
	}
}
