package general

import "fmt"

type Repository1 struct{}

func (r Repository1) GetData() string {
	return "Data from repository"
}

type Service1 struct {
	repo Repository1
}

// Constructor
func NewService1(repo Repository1) *Service1 {
	return &Service1{repo: repo}
}

func (s *Service1) Fetch() {
	fmt.Println(s.repo.GetData())
}

func dependecyInjectionConstructor() {
	repo := Repository1{}
	service := NewService1(repo) // Injecting repo
	service.Fetch()
}
