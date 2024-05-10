package infrastructure

import (
	"context"
	"smart/domain/model"
	"smart/domain/repository"
	"smart/infrastructure"
	"testing"
)

var tcases = []model.Issue{
	{Title: "Issue_1", Description: "Desc ....", Service: "scms", StoryPoint: "5"},
	{Title: "Issue_2", Description: "Desc ....", Service: "scms", StoryPoint: "5"},
	{Title: "Issue_3", Description: "Desc ....", Service: "scms", StoryPoint: "5"},
	{Title: "Issue_4", Description: "Desc ....", Service: "scms", StoryPoint: "5"},
}

func prepare(t *testing.T) repository.Issue {
	db, err := infrastructure.NewDBClient()
	if err != nil {
		t.Error(err)
	}
	return NewIssueRepo(db)
}

func clear(t *testing.T, repo repository.Issue) {
	for _, tc := range tcases {
		if err := repo.Delete(context.Background(), tc.ID); err != nil {
			t.Error(err)
		}
	}
}

func TestIssueRepo_Create(t *testing.T) {
	repo := prepare(t)
	defer clear(t, repo)

	for i := 0; i < len(tcases); i++ {
		if err := repo.Create(context.Background(), &tcases[i]); err != nil {
			t.Error(err)
		}
	}
}

func TestIssueRepo_List(t *testing.T) {
	repo := prepare(t)
	defer clear(t, repo)

	for i := 0; i < len(tcases); i++ {
		if err := repo.Create(context.Background(), &tcases[i]); err != nil {
			t.Error(err)
		}
	}

	models, err := repo.List(context.Background())
	if err != nil {
		t.Error(err)
	}

	if len(models) != len(tcases) {
		t.Error("length isn't compare")
	}
}
