package repository

import (
    "github.com/maguowei/example/internal/example/domain/entities"
    "github.com/maguowei/example/internal/example/domain/repository/ent"
)

type exampleRepository struct {
    entClient *ent.Client
}

func NewExampleRepository(entClient *ent.Client) ExampleRepositoryInterface {
    return &exampleRepository{
        entClient: entClient,
    }
}

func (r *exampleRepository) GetExample(exampleId int64) (*entities.Example, error) {
    panic("implement me")
}
func (r *exampleRepository) GetExamples() ([]entities.Example, error) {
    panic("implement me")
}
func (r *exampleRepository) CreateExample(example *entities.Example) (*entities.Example, error) {
    panic("implement me")
}

func (r *exampleRepository) UpdateExample(example *entities.Example) (*entities.Example, error) {
    panic("implement me")
}

func (r *exampleRepository) DeleteExample(exampleId int64) error {
    panic("implement me")
}
