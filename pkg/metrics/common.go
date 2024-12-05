package metrics

type Rankable interface {
	GetPackage() string
	GetName() string
	GetScore() int
}
