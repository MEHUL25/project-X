package general

import "fmt"

type Repository interface {
	GetData() string
}

type SQLRepository struct{}

func (SQLRepository) GetData() string {
	return "Data from SQL"
}

type APIService struct {
	repo Repository
}

func NewAPIService(repo Repository) *APIService {
	return &APIService{repo: repo}
}

func (a *APIService) Fetch() {
	fmt.Println(a.repo.GetData())
}

func main() {
	sqlRepo := SQLRepository{}
	apiService := NewAPIService(sqlRepo)
	apiService.Fetch()
}
