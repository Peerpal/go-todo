package repository

type Repository interface {
	Create(input interface{}) (interface{}, error)
	FindOne(id string) (interface{}, error)
}
