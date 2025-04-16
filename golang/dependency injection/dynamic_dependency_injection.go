package general

import (
	"fmt"
)

// Dependency Interface
type Logger1 interface {
	Log(message string)
}

// Concrete Implementations
type ConsoleLogger1 struct{}

func (c ConsoleLogger1) Log(message string) { fmt.Println("Console Log:", message) }

type FileLogger struct{}

func (f FileLogger) Log(message string) { fmt.Println("File Log:", message) }

// Service2 that depends on Logger1
type Service2 struct {
	logger Logger1
}

// Factory Function (Runtime DI)
func NewService2(loggerType string) *Service2 {
	var logger Logger1

	if loggerType == "console" {
		logger = ConsoleLogger1{}
	} else if loggerType == "file" {
		logger = FileLogger{}
	}

	return &Service2{logger: logger}
}

func dynamicDependencyInjectionExample() {
	service := NewService2("file") // Decides dependency at runtime
	service.logger.Log("Hello, Dynamic DI!")
}
