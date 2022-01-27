package repository

import (
	"github.com/maguowei/example/internal/example/domain/entities"
)

type ExampleRepositoryInterface interface {
	GetExample(exampleId int64) (*entities.Example, error)
	GetExamples() ([]entities.Example, error)
	CreateExample(example *entities.Example) (*entities.Example, error)
	UpdateExample(example *entities.Example) (*entities.Example, error)
	DeleteExample(exampleId int64) error
}