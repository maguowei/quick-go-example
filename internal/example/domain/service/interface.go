package service

import "github.com/maguowei/example/internal/example/domain/entities"

type ExampleDomainInterface interface {
    CreateExample(example *entities.Example) (*entities.Example, error)
}
