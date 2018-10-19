package github

import (
	"context"
	"fmt"
	"strings"

	hub "github.com/google/go-github/github"
	"github.com/sirupsen/logrus"
)

type GitHub struct {
	client *hub.Client
	ctx    context.Context
}

func New(username, password string) *GitHub {
	logrus.Info("Create GitHub client")
	// basic auth
	tp := hub.BasicAuthTransport{
		Username: strings.TrimSpace(username),
		Password: strings.TrimSpace(password),
	}
	client := hub.NewClient(tp.Client())
	ctx := context.Background()

	return &GitHub{
		client,
		ctx,
	}
}

func (g *GitHub) Release(org, repo, version, commit string) error {
	name := fmt.Sprintf("Release %s", version)
	newRelease := &hub.RepositoryRelease{
		TagName:         &version,
		TargetCommitish: &commit,
		Name:            &name,
		Draft:           hub.Bool(false),
		Prerelease:      hub.Bool(false),
	}
	rel, _, err := g.client.Repositories.CreateRelease(g.ctx, org, repo, newRelease)
	if err != nil {
		return fmt.Errorf("Error releasing repo %s to version %s\n %s", repo, version, err.Error())
	}

	logrus.Infof("Successfully released repo %s to version %s: %s", repo, version, rel.GetHTMLURL())
	return nil
}

func (g *GitHub) createPullRequest(org, repo, ciBranch, submodule, version, commit string) error {
	title := fmt.Sprintf("[DroneCI] update %s to version %s", submodule, version)
	branch := "master"
	description := fmt.Sprintf("Updated to version %s at commit %s", version, commit)
	newPR := &hub.NewPullRequest{
		Title:               &title,
		Head:                &ciBranch,
		Base:                &branch,
		Body:                &description,
		MaintainerCanModify: hub.Bool(true),
	}
	pr, _, err := g.client.PullRequests.Create(g.ctx, org, repo, newPR)
	if err != nil {
		return fmt.Errorf("Error while creating Pull Request for %s while updating submodule %s to version %s at commit %s: \n %s",
			repo, submodule, version, commit, err.Error())
	}

	logrus.Infof("Pull request %s for repo %s successfully created", pr.GetHTMLURL(), repo)
	return nil
}
