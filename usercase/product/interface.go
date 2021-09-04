package product

import(
	"github.com/pedrazzani/golang-file2api/entity"
)

type Repository interface {
	List() ([]entity.Product, error)
}