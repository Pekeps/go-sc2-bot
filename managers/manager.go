package managers

type manager interface {
	Init()
	Manage(loop uint32)
	Destruct()
}
