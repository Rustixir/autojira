package infrastructure

import (
	"context"
	"gorm.io/gorm"
	"smart/domain/model"
	"smart/domain/repository"
)

type issueRepo struct {
	db *gorm.DB
}

func NewIssueRepo(db *gorm.DB) repository.Issue {
	return issueRepo{db: db}
}

func (i issueRepo) Create(ctx context.Context, model *model.Issue) error {
	return i.db.Create(model).Error
}

func (i issueRepo) List(ctx context.Context) ([]model.Issue, error) {
	var list []model.Issue
	err := i.db.Find(&list).Error
	return list, err
}

func (i issueRepo) Delete(ctx context.Context, id uint) error {
	return i.db.Delete(&model.Issue{}, id).Error
}
