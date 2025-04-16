package general

import "fmt"

// Dependency (Interface)
type Logger interface {
	Log(message string)
}

// Concrete Implementation
type ConsoleLogger struct{}

func (c ConsoleLogger) Log(message string) {
	fmt.Println("Log:", message)
}

// Service that depends on Logger
type Service struct {
	logger Logger // Injected dependency
}

// Constructor for Service
func NewService(logger Logger) *Service {
	return &Service{logger: logger}
}

func (s *Service) DoSomething() {
	s.logger.Log("Service is doing something")
}

func dependencyInjectionExample() {
	logger := ConsoleLogger{}     // Creating dependency
	service := NewService(logger) // Injecting dependency

	service.DoSomething()
}

// Service depends on Logger, but doesn't create it.
// We inject a ConsoleLogger into Service.
