package service

import (
	"errors"
	"fmt"
	"github.com/andygrunwald/go-jira"
	"smart/domain/model"
	"smart/pkg/errhandler"
)

type JiraService interface {
	RegisterIssue(issue *model.Issue) error
}

type jiraService struct {
	client *jira.Client
}

func NewJiraService(client *jira.Client) JiraService {
	return jiraService{client}
}

func (j jiraService) RegisterIssue(issue *model.Issue) error {
	resp, _, err := j.client.Issue.Create(&jira.Issue{
		ID: fmt.Sprintf("ai-%s-%d", issue.Project, issue.ID),
		Fields: &jira.IssueFields{
			Description: issue.Description,
			Summary:     issue.Summary,
			Type: jira.IssueType{
				Name: issue.Type,
			},
			Project: jira.Project{
				Key: issue.Project,
			},
		},
	})

	if resp == nil {
		return errors.New("expected issue. Issue is nil")
	}
	if err != nil {
		return errhandler.Wrap("error given", err)
	}

	return nil
}
