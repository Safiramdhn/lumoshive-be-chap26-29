package repositories

type MainRepo[T any] interface {
	Create(entityInput T) (T, error)
}
