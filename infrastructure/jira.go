package infrastructure

import "github.com/andygrunwald/go-jira"

func NewJira(url string, user string, pass string) (*jira.Client, error) {
	tp := jira.BasicAuthTransport{
		Username: user,
		Password: pass,
	}

	return jira.NewClient(tp.Client(), url)
}
