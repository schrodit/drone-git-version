package main

import (
	"github.com/schrodit/drone-git-version/git"
	"github.com/schrodit/drone-git-version/version"
)

type (
	// Config maps the params we need to run Helm
	Config struct {
		GitName        string `json:"git_name"`
		GitEmail       string `json:"git_email"`
		InputFile      string `json:"input_file"`
		OutputFile     string `json:"output_file"`
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

	version.UpdateVersionFile(p.Config.InputFile, p.Config.OutputFile)
	Git.Commit(p.Config.OutputFile)
	Git.Push(p.Config.Branch)

	return nil
}
