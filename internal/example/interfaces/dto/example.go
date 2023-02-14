package dto

import "example/internal/example/domain/entities"

type ExampleCreateReqDto struct {
	//Title    string `json:"title" binding:"required"`
}

type ExampleCreateRespDto struct {
	ID int64 `json:"id"`
}

func (dto *ExampleCreateReqDto) ToDO() *entities.Example {
	return &entities.Example{}
}
