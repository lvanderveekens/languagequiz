package drill

type Storage interface {
	CreateDrill(c CreateDrillCommand) (*Drill, error)
}
