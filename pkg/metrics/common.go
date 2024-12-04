package metrics

type Rankable interface {
	GetName() string
	GetScore() int
}
