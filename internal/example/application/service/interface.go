package service

import (
    "context"
    "example/internal/example/domain/entities"
)

type ExampleAppServiceInterface interface {
    CreateExample(ctx context.Context, example *entities.Example) (*entities.Example, error)
}
