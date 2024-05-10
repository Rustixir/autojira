package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"smart/application/service"
	"smart/infrastructure"
	infrastructure2 "smart/infrastructure/repository"
	"smart/presenter/http"
)

func main() {

	log.Println("Initialize ...")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url := os.Getenv("JIRA_URL")
	user := os.Getenv("JIRA_USER")
	pass := os.Getenv("JIRA_PASS")
	token := os.Getenv("OPENAI_TOKEN")

	db, err := infrastructure.NewDBClient()
	if err != nil {
		log.Fatal(err)
	}

	jira, _ := infrastructure.NewJira(url, user, pass)
	ai := infrastructure.NewOpenAI(token)

	issueRepo := infrastructure2.NewIssueRepo(db)

	aiSvc := service.NewAiService(ai)
	jiraSvc := service.NewJiraService(jira)
	svc := service.NewIssueService(issueRepo, aiSvc, jiraSvc)
	http.Bootstrap(svc)
}
