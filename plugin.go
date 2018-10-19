package main

import (
	"github.com/schrodit/drone-git-version/git"
	"github.com/schrodit/drone-git-version/github"
	"github.com/schrodit/drone-git-version/versions"
)

type (
	// Config maps the params we need to run Helm
	Config struct {
		GitHubUsername string `json:"github_username"`
		GitHubPassword string `json:"github_password"`
		GitName        string `json:"git_name"`
		GitEmail       string `json:"git_email"`
		File           string `json:"file"`
		RepoOwner      string `json:"repo_owner"`
		RepoName       string `json:"repo_name"`
		Branch         string `json:"branch"`
		DeploymentType string `json:"deployment_type"`
	}
	// Plugin default
	Plugin struct {
		Config  Config
		command []string
	}
)

func (p *Plugin) Exec() error {
	Git := git.New("./", p.Config.GitName, p.Config.GitEmail)
	Github := github.New(p.Config.GitHubUsername, p.Config.GitHubPassword)

	newVersion := versions.UpdateVersionFile(p.Config.File, p.Config.DeploymentType)
	hash := Git.Commit(p.Config.File, newVersion)
	Git.Push(p.Config.GitHubUsername, p.Config.GitHubPassword, hash, p.Config.Branch)
	// TODO: make GitHub release and tag commit
	Github.Release(p.Config.RepoOwner, p.Config.RepoName, newVersion, hash)

	// TODO: make pull request wesense_landscape to upgrade current submodule repo to tagged release

	return nil
}
