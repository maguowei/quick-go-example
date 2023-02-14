package service

import (
    "context"
    "example/internal/example/domain/entities"
    "example/internal/example/domain/repository"
)

type exampleDomainService struct {
    exampleRepo repository.ExampleRepositoryInterface
}

func NewExampleDomainService(repository repository.ExampleRepositoryInterface) ExampleDomainInterface {
    return &exampleDomainService{
        exampleRepo: repository,
    }
}

func (s *exampleDomainService) CreateExample(ctx context.Context, example *entities.Example) (*entities.Example, error) {
    return s.exampleRepo.CreateExample(ctx, example)
}
