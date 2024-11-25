package backend

type Factory interface {
	// Type
	Type() Type

	// Reader
	Reader(*Config) (Backend, error)
}
