package repository

import(
	"os"
	"bufio"
	"strconv"

	"github.com/pedrazzani/golang-file2api/entity"
)

type FileRepository struct {
	path string
}

func NewFileRepository(path string) *FileRepository {
	return &FileRepository{
		path: path,
	}
}

func (r *FileRepository) List() ([]entity.Product, error) {
	f, err := os.Open(string(r.path))
    
    scanner := bufio.NewScanner(f)

    products := []entity.Product{}

    for scanner.Scan() {
        
        valueInCents, _ := strconv.Atoi(string(scanner.Bytes()[11:17]))
        value := float32(valueInCents)/100
        days, _ := strconv.Atoi(string(scanner.Bytes()[17:20]))
        product := entity.NewProduct(
                string(scanner.Bytes()[:2]),
                string(scanner.Bytes()[2:4]),
                string(scanner.Bytes()[4:5]),
                string(scanner.Bytes()[5:11]),
                value,
                days,
                string(scanner.Bytes()[20:46]),
                string(scanner.Bytes()[46:72]))

        products = append(products, *product)
    }

    f.Close()

    return products, err
}