package service

import "github.com/maguowei/example/internal/example/domain/entities"

type ExampleAppServiceInterface interface {
    CreateExample(example *entities.Example) (*entities.Example, error)
}
