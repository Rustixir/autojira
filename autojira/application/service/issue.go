package service

import (
	"context"
	"smart/domain/model"
	"smart/domain/repository"
)

type IssueService interface {
	Create(ctx context.Context, issueMessage string) error
	List(ctx context.Context) ([]model.Issue, error)
	Delete(ctx context.Context, id uint) error
}

type issueService struct {
	AiService
	JiraService
	issueRepo repository.Issue
}

func NewIssueService(issueRepo repository.Issue, ai AiService, jira JiraService) IssueService {
	return issueService{
		ai,
		jira,
		issueRepo,
	}
}

func (i issueService) Create(ctx context.Context, issueMessage string) error {
	issue, err := i.ProcessIssue(issueMessage)
	if err != nil {
		return err
	}

	if err = i.issueRepo.Create(ctx, issue); err != nil {
		return err
	}

	return i.RegisterIssue(issue)
}

func (i issueService) List(ctx context.Context) ([]model.Issue, error) {
	//TODO implement me
	panic("implement me")
}

func (i issueService) Delete(ctx context.Context, id uint) error {
	//TODO implement me
	panic("implement me")
}
