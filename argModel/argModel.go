package argModel

type Entity interface {
	New() Entity
	GetID() string
	SetID(string)
}

func GetEntity(model string) Entity {
	switch model {
	case "user":
		return &User{}
	default:
		return nil
	}
}
