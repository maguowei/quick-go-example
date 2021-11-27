package repository

import (
	"github.com/maguowei/example/internal/example/domain/models"
)

type ExampleRepository interface {
	GetExample(exampleId int64) (models.Example, error)
	GetExamples() ([]models.Example, error)
	CreateExample(example models.Example) (models.Example, error)
	UpdateExample(example models.Example) (models.Example, error)
	DeleteExample(exampleId int64) error
}