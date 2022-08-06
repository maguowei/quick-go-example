package repository

import (
	"context"
	"github.com/maguowei/example/internal/example/domain/entities"
)

type ExampleRepositoryInterface interface {
	GetExample(ctx context.Context, exampleId int64) (*entities.Example, error)
	GetExamples(ctx context.Context, exampleIds []int64) ([]entities.Example, error)
	CreateExample(ctx context.Context, example *entities.Example) (*entities.Example, error)
	UpdateExample(ctx context.Context, example *entities.Example) (*entities.Example, error)
	DeleteExample(ctx context.Context, exampleId int64) error
}