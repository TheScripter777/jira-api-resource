package main

import (
	"fmt"
	"github.com/TurnsCoffeeIntoScripts/jira-api-issue-resource/pkg/application"
	"os"
)

func main() {
	app := &application.JiraAPIResourceApp{}
	if err := app.Run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
