package repository

type Repository interface {
	Create(input interface{}) (interface{}, error)
	FindMany() (interface{}, error)
	FindOne(id string) (interface{}, error)
}
