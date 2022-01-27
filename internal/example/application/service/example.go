package service

import (
    "github.com/maguowei/example/internal/example/domain/entities"
    "github.com/maguowei/example/internal/example/domain/service"
)

type exampleAppService struct {
    exampleDomainService service.ExampleDomainInterface
}

func NewExampleAppService(exampleDomainService service.ExampleDomainInterface) ExampleAppServiceInterface {
    return &exampleAppService{
        exampleDomainService: exampleDomainService,
    }
}

func (s *exampleAppService) CreateExample(example *entities.Example) (*entities.Example, error) {
    return s.exampleDomainService.CreateExample(example)
}
