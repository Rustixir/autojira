package repository

import (
	"context"
	"smart/domain/model"
)

type Issue interface {
	Create(ctx context.Context, model *model.Issue) error
	List(ctx context.Context) ([]model.Issue, error)
	Delete(ctx context.Context, id uint) error
}
