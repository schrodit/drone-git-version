package main

import (
	"github.com/schrodit/drone-git-version/git"
	"github.com/schrodit/drone-git-version/versions"
)

type (
	// Config maps the params we need to run Helm
	Config struct {
		GitName        string `json:"git_name"`
		GitEmail       string `json:"git_email"`
		File           string `json:"file"`
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

	newVersion := versions.UpdateVersionFile(p.Config.File, p.Config.DeploymentType)
	Git.Commit(p.Config.File, newVersion)
	Git.Push(p.Config.Branch)
	// TODO: make GitHub release and tag commit

	// TODO: make pull request wesense_landscape to upgrade current submodule repo to tagged release

	return nil
}
