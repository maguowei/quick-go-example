package repository

import (
    "context"
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

func (r *exampleRepository) GetExample(ctx context.Context, exampleId int64) (*entities.Example, error) {
    panic("implement me")
}
func (r *exampleRepository) GetExamples(ctx context.Context, exampleIds []int64) ([]entities.Example, error) {
    panic("implement me")
}
func (r *exampleRepository) CreateExample(ctx context.Context, example *entities.Example) (*entities.Example, error) {
    panic("implement me")
}

func (r *exampleRepository) UpdateExample(ctx context.Context, example *entities.Example) (*entities.Example, error) {
    panic("implement me")
}

func (r *exampleRepository) DeleteExample(ctx context.Context, exampleId int64) error {
    panic("implement me")
}
