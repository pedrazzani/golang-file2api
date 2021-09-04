package product

import(
    "github.com/pedrazzani/golang-file2api/entity"
)

type Service struct {
    repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (service *Service) List() ([]entity.Product, error) {
    return service.repository.List()
}