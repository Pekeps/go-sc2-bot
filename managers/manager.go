package managers

import "github.com/pekeps/go-sc2ai/builds"

type Manager interface {
	Init()
	Manage()

	GetBuildRequests() []*builds.BuildRequest
}
