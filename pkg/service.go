package pkg

import "fmt"

type DemoService interface {
	Demo(string) (string, error)
}

type demoService struct {
}

func (demoService) Demo(s string) (string, error) {
	return fmt.Sprintf("Hello %s!", s), nil
}

func NewDemoService() DemoService {
	return &demoService{}
}
