package service

import (
    "context"
    "example/internal/example/domain/entities"
    "example/internal/example/domain/service"
)

type exampleAppService struct {
    exampleDomainService service.ExampleDomainInterface
}

func NewExampleAppService(exampleDomainService service.ExampleDomainInterface) ExampleAppServiceInterface {
    return &exampleAppService{
        exampleDomainService: exampleDomainService,
    }
}

func (s *exampleAppService) CreateExample(ctx context.Context, example *entities.Example) (*entities.Example, error) {
    return s.exampleDomainService.CreateExample(ctx, example)
}
